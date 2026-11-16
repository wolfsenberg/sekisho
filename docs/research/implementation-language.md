# Implementation language for Sekisho

Research note, 2026-11-16. Question: what is the best and most optimized implementation language for Sekisho, and should it stay in Go?

All vendor benchmarks below are **self-reported by the vendor of the winning gateway**. None has been independently reproduced for this note. Items that could not be verified against a primary source are marked *unverified*.

## 1. TL;DR

- **Stay in Go.** Sekisho's request latency is dominated by upstream LLM time-to-first-token and generation (hundreds of ms to minutes). Published gateway overheads for compiled-language gateways (Go, Rust) are in the **microsecond to low-millisecond range** ([Bifrost](https://docs.getbifrost.ai/benchmarking/getting-started), [TensorZero](https://github.com/tensorzero/tensorzero/blob/main/docs/gateway/benchmarks.mdx), [Helicone](https://github.com/Helicone/ai-gateway/blob/main/benchmarks/README.md)). That is 2 to 5 orders of magnitude below the upstream latency.
- Rust gateways do publish lower tail overhead and much lower memory than Go gateways (for example, LiteLLM's own July 2026 bench: Rust p99 ~0.7 ms / 21.8 MB vs Bifrost (Go) 4.5 ms / 199 MB; [LiteLLM blog](https://docs.litellm.ai/blog/rust-ai-gateway-benchmarks)). But neither number matters for a self-hosted, multi-workspace gateway whose traffic is in the tens to hundreds of concurrent streams.
- The decisive factor is **migration cost**: Sekisho's provider executors and translators come from the Go library CLIProxyAPI v8 (MIT, Go, ~58.7k stars, very active; [repo](https://github.com/router-for-me/CLIProxyAPI)). No other language has an equivalent library covering Codex/Claude Code/Gemini CLI/Antigravity/Grok subscription OAuth flows. Rewriting means owning all provider-protocol churn yourself.
- Python (LiteLLM) and, to a lesser degree, Node are the only options with measurable overhead problems under load in published benchmarks. Zig is pre-1.0 (0.17.1; [ziglang.org](https://ziglang.org/download/)), so it is unsuitable.
- Do the Go-side optimizations in section 7 first (transport pool sizing, GOMEMLIMIT, streaming without buffering, one HTTP router, pprof-driven work).

## 2. What Sekisho's workload actually is

Verified in the repo:

- Single Go binary, `CGO_ENABLED=0`, cross-compiled per `GOOS/GOARCH` (`Makefile` lines 26, 68; `Dockerfile` line 25). The React dashboard is embedded.
- Go 1.26, chi v5 **and** gin v1.10, gorilla/websocket, go-webauthn, pquerna/otp, filippo.io/age, modernc.org/sqlite (pure-Go SQLite) (`go.mod`).
- 438 `.go` files, ~74.4k lines. CLIProxyAPI v8.0.35 is imported in **13 files, all in `internal/upstream/`**. (The task brief said 17. I counted 13 `.go` files with 26 import lines across `sdk/translator`, `sdk/cliproxy/executor`, `sdk/cliproxy/auth`, `sdk/auth`, `sdk/config`, `sdk/api`.)
- The upstream client clones `http.DefaultTransport` with `MaxIdleConns=16`, `MaxIdleConnsPerHost=8`, and a 5-minute inference header timeout (`internal/upstream/client.go` ~L100-125).

Workload shape:

- **I/O-bound proxying** of HTTP, SSE and WebSocket streams. The CPU work per request is JSON translation between protocol dialects (OpenAI Responses/Chat, Claude Messages, Gemini), token/usage accounting, routing decisions, and a few SQLite writes.
- **Long-lived connections.** Agentic coding sessions (Codex, Claude Code) hold streams for seconds to minutes. Concurrency is roughly users times parallel sessions, not RPS.
- **Latency budget.** Upstream TTFT is typically 100s of ms to several seconds. In TensorZero's benchmark, even LiteLLM at 100 QPS added only 5.87 ms p99 ([source](https://github.com/tensorzero/tensorzero/blob/main/docs/gateway/benchmarks.mdx)). The agentgateway authors and the LiteLLM authors both concede that proxy overhead "matters less" against model latency ([agentgateway](https://agentgateway.dev/blog/2026-06-26-benchmarking-agentgateway-vs-litellm-part-2/), [LiteLLM](https://docs.litellm.ai/blog/rust-ai-gateway-benchmarks)).
- **Where overhead does compound:** multi-turn agent loops (30 turns). In LiteLLM's bench, added wall time per 30-turn session was Rust 0.03 s, Bifrost (Go) 0.13 s, Portkey (TS) 0.12 s, and Python 0.97 s ([source](https://docs.litellm.ai/blog/rust-ai-gateway-benchmarks)). Against a session that runs for minutes, even 0.13 s is noise.
- **Deployment:** self-hosted, single node, SQLite. Memory and binary size matter more than peak RPS.

## 3. Language comparison

Ratings: ++ excellent, + good, 0 adequate, - weak, -- poor, for Sekisho specifically.

| Criterion | Go | Rust | Zig | C++ | Java/Kotlin | C#/.NET AOT | Elixir/BEAM | TS (Node/Bun) | Python |
|---|---|---|---|---|---|---|---|---|---|
| Proxy throughput / added latency | + (µs to low ms) | ++ (sub-ms p99) | + (in theory) | ++ | + | + | 0 | 0 (Portkey 2.3 ms p99) | -- (LiteLLM 257 ms p99) |
| Memory idle/load | + (10s to 100s MB) | ++ (13 to 100 MB published) | ++ | ++ | - JIT / 0 native-image | + AOT | 0 | 0 | -- |
| Many long-lived SSE/WS | ++ goroutines | ++ tokio | 0 (immature async) | + (hard) | + virtual threads | + async/await | ++ processes | + event loop | 0 asyncio |
| GC vs tail latency | + concurrent GC, short STW | ++ no GC | ++ no GC | ++ no GC | + (G1/ZGC; native-image Serial) | + | + per-process GC | 0 | - |
| Ecosystem (OAuth2, WebAuthn, age, SQLite, LLM protocols) | ++ (all in use today, plus CLIProxyAPI) | + (age crate is "[BETA]", webauthn-rs 0.6) | -- | - | + (no age impl verified) | + (no age impl verified) | 0 | + | ++ for LLM SDKs |
| Single binary + cross-compile | ++ (`CGO_ENABLED=0`) | + (needs C toolchains for some crates) | ++ | - | - native-image (no cross-compile verified) | - (no cross-OS) | - (release dir) | 0 Bun `--compile --target`; Node SEA "active development" | -- |
| Dev speed under API churn | + | 0 | -- | -- | + | + | 0 | ++ | ++ |
| Migration cost from current code | none | very high | very high | very high | high | high | very high | high | high |

Notes on the table: latency figures come from the gateway benchmarks in section 4. Runtime facts come from section 5 sources.

## 4. Real-world gateways

| Gateway | Language | Published numbers | Who ran it, hardware, method | Source |
|---|---|---|---|---|
| **CLIProxyAPI** (Sekisho's dependency) | Go (`go.mod`, "Reusable Go SDK"); README mentions a separate Rust port (*unverified*) | None published | n/a | [repo](https://github.com/router-for-me/CLIProxyAPI) |
| **one-api** | Go (single binary, SQLite default) | None. Only config defaults (e.g. `SQL_MAX_OPEN_CONNS` 1000) | n/a | [repo](https://github.com/songquanpeng/one-api) |
| **new-api** | Go + Gin backend, React 19 frontend embedded. Fork of one-api | None | n/a | [repo](https://github.com/QuantumNous/new-api) |
| **Bifrost** (Maxim) | Go | 5k RPS sustained: overhead 59 µs (t3.medium), 11 µs (t3.xlarge). Peak memory 1,313 MB / 3,340 MB. vs LiteLLM at 500 RPS on t3.medium: 424 vs 44.8 req/s achieved, p99 1.68 s vs 90.7 s, peak mem 120 MB vs 372 MB. Overhead with 60 ms mock: 0.99 ms vs 40 ms | Maxim (vendor). AWS t3.medium (2 vCPU/4 GB) and t3.xlarge (4 vCPU/16 GB), mocked OpenAI. **Response sizes differ between the two instance runs (~1 KB vs ~10 KB)**, which confounds the comparison. Large buffer and pool presizing (15k to 20k) explains the high memory | [docs](https://docs.getbifrost.ai/benchmarking/getting-started), [Maxim page](https://www.getmaxim.ai/bifrost/resources/benchmarks.md), [repo](https://github.com/maximhq/bifrost) |
| **TensorZero** | Rust | Overhead at 10k QPS: mean 0.37 ms, p50 0.35, p99 0.94 ms. LiteLLM at 100 QPS: p99 5.87 ms. At 500 QPS: p99 39.69 ms. At 1,000 QPS: failure | TensorZero (vendor). One AWS c7i.xlarge (4 vCPU/8 GB), load generator + gateway + mock on the **same host**, TensorZero observability disabled. Run 2025-07-30, TZ 2025.5.7 vs LiteLLM 1.74.9. Repo shows as archived (June 2026) per GitHub page | [benchmarks.mdx](https://github.com/tensorzero/tensorzero/blob/main/docs/gateway/benchmarks.mdx), [repo](https://github.com/tensorzero/tensorzero) |
| **Helicone AI Gateway** | Rust | 3,000 RPS sustained, overhead "<1 ms", p95 89.16 ms incl. 60 ms mock, memory <100 MB, ~80% CPU. README table: ~64 MB memory, ~30 MB binary vs a vague "typical setup" | Helicone (vendor). Fly.io performance-2x with 4 GB (AMD EPYC), k6, mock with 60 ms median, 2-minute run, logging disabled. The "typical setup" comparison column is undefined | [benchmarks README](https://github.com/Helicone/ai-gateway/blob/main/benchmarks/README.md), [repo](https://github.com/Helicone/ai-gateway) |
| **agentgateway** (Linux Foundation) | Rust (`Cargo.toml`) | Fixed 3,000 QPS: p50 0.227 ms, p99 0.436 ms, ~13 MiB avg mem, 13.4% CPU. LiteLLM (18 workers): p50 12.3 ms, p99 30.6 ms, 2,466 QPS achieved, 11.67 GiB mem, 345% CPU | agentgateway project (vendor). Fortio, 32 connections, 1 KB payloads, 30 s, mock server. Hardware **not stated** | [blog](https://agentgateway.dev/blog/2026-06-26-benchmarking-agentgateway-vs-litellm-part-2/), [repo](https://github.com/agentgateway/agentgateway) |
| **LiteLLM proxy** | Python | 4 instances: 1,170 RPS, overhead p50 2 ms / p95 8 ms / p99 13 ms. 2 instances: p99 43 ms. 33 pods: 3,000 RPS | LiteLLM (vendor). 4 vCPU/8 GB per instance, 4 workers, Locust 1,000 users, mock endpoint, Postgres, Redis off. Overhead from `x-litellm-overhead-duration-ms` header (self-measured inside the process). No date | [benchmarks](https://docs.litellm.ai/docs/benchmarks) |
| **LiteLLM Rust gateway** (early beta) | Rust | p99 added: Rust 0.7 ms, Portkey 2.3 ms, Bifrost 4.5 ms, LiteLLM Python 257.7 ms. Peak memory: 21.8 / 90.4 / 199.1 / 329.5 MB. Throughput: Rust ~2,814 vs Bifrost ~2,744 req/s | LiteLLM (vendor, benchmarking its own new product against competitors). Single host, **no repeated trials or error bars**. The authors say to read results as order-of-magnitude. Hardware not stated beyond a 4 vCPU/16 GB cost assumption. 2026-07-22 | [blog](https://docs.litellm.ai/blog/rust-ai-gateway-benchmarks) |
| **Portkey gateway** | TypeScript (Node/Workers; `package.json`, `wrangler.toml`) | "<1 ms latency", "122 kb" footprint | Portkey (vendor). **No methodology published** | [repo](https://github.com/Portkey-AI/gateway) |
| **Kong AI Gateway** | Built on Kong Gateway (Nginx/OpenResty + Lua is generally known; *not stated on the benchmark page*) | Relative only: "228% faster than Portkey, 859% faster than LiteLLM" in RPS, "65%/86% lower latency". Absolute numbers only in charts | Kong (vendor). EKS 1.32, c5.4xlarge nodes, each gateway capped at 12 CPUs, k6 400 VUs, WireMock, 3-minute runs, proxy-only. Kong 3.10, Portkey 1.9.19, LiteLLM 1.63.7 | [Kong blog](https://konghq.com/blog/engineering/ai-gateway-benchmark-kong-ai-gateway-portkey-litellm) |
| **Envoy AI Gateway** (now "Agent Router") | Go control plane (`go.mod`) + Envoy (C++) data plane | None found | n/a | [repo](https://github.com/envoyproxy/ai-gateway), [docs](https://theagentrouter.ai/docs/) |

Critical reading:

- **Every benchmark is vendor-run, against a mock upstream, usually with logging and persistence disabled.** None includes real provider TLS, real streaming durations, or database writes per request, which are Sekisho's real costs.
- The benchmarks disagree with each other. Bifrost reports 11 to 59 µs overhead for itself, while LiteLLM's bench reports 4.5 ms p99 for Bifrost. LiteLLM's own doc reports 2 ms p50 overhead for its Python proxy, while agentgateway measured 12.3 ms p50 for it. Treat any single number as ±10x.
- **Methodology problems:** TensorZero ran the load generator on the same 4 vCPU host. Bifrost changed the payload size between instance types. LiteLLM self-reports overhead from a header its own process sets. Kong publishes only relative percentages.
- What survives across all sources: **(a)** Python is the outlier, from 10s to 100s of ms tail and GBs of memory under load; **(b)** Rust gateways show the lowest memory (13 to 100 MB) and sub-ms p99; **(c)** Go gateways sit within a few ms of Rust; **(d)** Go is the dominant language among the subscription/key-pool gateways most similar to Sekisho (one-api, new-api, CLIProxyAPI).

## 5. Per-language notes

**Go (current).** Goroutines fit one-goroutine-per-stream proxying. The GC is mostly concurrent with "brief stop-the-world pauses" between phases ([GC guide](https://tip.golang.org/doc/gc-guide)). Go 1.26 enables the Green Tea GC by default, which Go expects to cut GC overhead by 10 to 40% in GC-heavy programs ([Go 1.26 notes](https://go.dev/doc/go1.26)). `GOMEMLIMIT` gives a soft memory cap ([GC guide](https://tip.golang.org/doc/gc-guide)). Pure-Go deps (modernc SQLite, age) keep `CGO_ENABLED=0` cross-compilation trivial, and the repo already does this. The ecosystem covers every Sekisho need, including the only mature subscription-OAuth provider library (CLIProxyAPI).

**Rust.** Best published overhead and memory numbers (TensorZero, Helicone, agentgateway, LiteLLM-Rust). Libraries exist: `age` crate 0.12.1, but its description says "[BETA]" ([crates.io](https://crates.io/api/v1/crates/age)); `webauthn-rs` 0.6.2 ([crates.io](https://crates.io/api/v1/crates/webauthn-rs)); SQLite through rusqlite (*not fetched*). The costs are slower iteration when provider payloads change weekly, a rewrite of ~74k lines, and reimplementing CLIProxyAPI's executors. This is the right choice only if a hot path is ever proven to be CPU-bound (section 6).

**Zig.** Latest stable 0.17.1 (2026-11-02), no 1.0 ([ziglang.org](https://ziglang.org/download/)). No mature WebAuthn, age, or OAuth libraries were verified. Not suitable.

**C++.** Envoy proves C++ data planes work at scale, but Envoy AI Gateway pairs it with a Go control plane ([repo](https://github.com/envoyproxy/ai-gateway)). Memory-safety risk for a credential vault, slow development, and difficult cross-compilation. Not suitable.

**Java/Kotlin.** Virtual threads "exist to provide scale (higher throughput), not speed (lower latency)" and pin carrier threads inside `synchronized` or native calls ([Oracle docs](https://docs.oracle.com/en/java/javase/21/core/virtual-threads.html)). Later JDKs address `synchronized` pinning (JEP 491; openjdk.org returned 403, so this is *unverified here*). GraalVM native-image defaults to the Serial GC, and G1 is Linux-only and "not available in GraalVM Community Edition" ([GraalVM docs](https://www.graalvm.org/latest/reference-manual/native-image/optimizations-and-performance/MemoryManagement/)). That is a tail-latency and single-binary trade-off. No gain over Go.

**C#/.NET Native AOT.** Self-contained native binary with smaller memory ([MS Learn](https://learn.microsoft.com/en-us/dotnet/core/deploying/native-aot/)). However, "Native AOT does not support cross-OS compilation" ([MS Learn](https://learn.microsoft.com/en-us/dotnet/core/deploying/native-aot/cross-compile)), it requires trimming, and it has no reflection emit. That is a worse release story than `GOOS=... go build`.

**Elixir/BEAM.** Excellent fit for millions of long-lived connections with per-process GC. But no single static binary (releases bundle the ERTS), no published LLM gateway numbers found, and a thin ecosystem for age and WebAuthn (*not verified*).

**TypeScript (Node/Bun).** Fastest iteration on provider JSON changes, and the dashboard already uses TS. Bun can build single-file executables and cross-compile with `--target` ([Bun docs](https://bun.com/docs/compiler/executables)). Node SEA is "Stability: 1.1 - Active development" and loads native addons only through a temp-file workaround ([Node docs](https://nodejs.org/api/single-executable-applications.html)). That affects SQLite. Portkey (TS) measured 2.3 ms p99 / 90 MB in LiteLLM's bench, which is acceptable but no better than Go.

**Python.** The worst published tail latency and memory under load (LiteLLM 257.7 ms p99 in LiteLLM's own bench; 11.67 GiB in agentgateway's bench). No single-binary story. LiteLLM itself is building a Rust gateway ([blog](https://docs.litellm.ai/blog/rust-ai-gateway-benchmarks)). Not suitable.

## 6. Recommendation for Sekisho

**Keep Go. Do not rewrite.** Reasoning:

1. Overhead is not the bottleneck. Published Go overhead is 0.01 to 5 ms. Upstream LLM latency is 100 to 10,000+ ms.
2. Memory at Sekisho's scale is fine. Go gateways report 100 to 200 MB under synthetic thousands of RPS (Bifrost 120 MB at 500 RPS). A self-hosted Sekisho instance will see far fewer requests.
3. CLIProxyAPI is a strategic asset. It absorbs provider-protocol churn upstream, and it exists only in Go. Every other language means maintaining Codex/Claude/Gemini CLI/Antigravity/Grok executors yourself.
4. Deployment already meets the target: static `CGO_ENABLED=0` binary, embedded UI, trivial cross-compilation.

**When the answer would change:**

| Condition | Changed answer |
|---|---|
| Sustained >5k RPS per node with short, non-streaming responses (for example, embeddings or classification traffic) and profiling shows gateway CPU or GC as the bottleneck | Consider a Rust data-plane sidecar for the hot path only. Keep the Go control plane, as Envoy AI Gateway and agentgateway do |
| Target hardware ≤512 MB RAM (Raspberry Pi, small VPS) with many tenants | Tune `GOMEMLIMIT` first. A Rust core becomes attractive only if Go cannot stay within budget |
| CLIProxyAPI is abandoned, relicensed, or diverges from Sekisho's needs | Fork it in Go. Rewrite cost stays the same in any other language, so this is still not a reason to switch |
| Team is primarily TS with no Go expertise | Bun/TS is viable for development speed, but it would still have to replace CLIProxyAPI. Go stays cheaper |
| Need for millions of idle WebSocket connections (push notifications) | BEAM or Rust. Not Sekisho's workload |

## 7. Go-side optimizations worth doing first

Measure before and after each change with a mock upstream (k6 or Fortio), `pprof`, and `runtime/metrics`.

1. **Size the upstream connection pool for streaming.** `internal/upstream/client.go` sets `MaxIdleConnsPerHost=8` and `MaxIdleConns=16`. Long SSE streams hold connections, so bursts above 8 per provider host open new TLS connections (handshake cost on every burst). Raise these per provider (for example 64 to 256), set `IdleConnTimeout`, and confirm HTTP/2 is negotiated. Go 1.26 adds `HTTP2Config.StrictMaxConcurrentRequests` to control opening new connections when the stream limit is reached ([Go 1.26](https://go.dev/doc/go1.26)).
2. **Set `GOMEMLIMIT`** (for example 70 to 80% of the container limit) and, if memory allows, raise `GOGC`. "Doubling GOGC will double heap memory overheads and roughly halve GC CPU cost" ([GC guide](https://tip.golang.org/doc/gc-guide)). Expose both as config and document them for self-hosters.
3. **Stay on Go 1.26+ with Green Tea GC** (default), which Go estimates at 10 to 40% less GC overhead ([Go 1.26](https://go.dev/doc/go1.26)).
4. **Stream, do not buffer.** `io.ReadAll` appears in `internal/server/gateway.go:215` and `internal/upstream/client.go:428` (bounded with `LimitReader`). Make sure SSE and WS response paths decode incrementally and flush per event. Reserve full buffering for request bodies that need translation.
5. **Reduce per-chunk allocation in translation.** SSE chunk to JSON to re-encode is the main CPU cost. Profile allocations (`-memprofile`) in the CLIProxyAPI translators, then reuse buffers with `sync.Pool`. Consider a faster JSON codec only if the profile shows encoding/json dominating (*no measurement yet*).
6. **Pick one HTTP router.** `go.mod` carries both chi and gin. Standardizing on chi (or `net/http` `ServeMux`) cuts binary size and dependency surface. Expect little latency gain.
7. **Batch SQLite writes for usage accounting** off the request path (buffered channel and a periodic transaction), as one-api does with `BATCH_UPDATE_INTERVAL` ([one-api](https://github.com/songquanpeng/one-api)). Use WAL mode and a single writer connection with modernc SQLite.
8. **Add a repeatable overhead benchmark** in `tools/` (mock upstream with 60 ms TTFT plus streamed chunks) measuring gateway-added p50/p99 and RSS. That gives Sekisho its own numbers instead of relying on vendor benchmarks.
9. **Reconsider a Rust hot path only if** steps 1 to 8 are done and profiles still show the gateway (not upstream) dominating latency or memory at the target scale.

## 8. Sources

Repo (read 2026-11-16): `go.mod`, `PRODUCT.md` (Stack), `Makefile`, `Dockerfile`, `internal/upstream/client.go`, `internal/server/gateway.go`.

Gateways:
- CLIProxyAPI: https://github.com/router-for-me/CLIProxyAPI
- one-api: https://github.com/songquanpeng/one-api
- new-api: https://github.com/QuantumNous/new-api
- Bifrost: https://github.com/maximhq/bifrost, https://docs.getbifrost.ai/benchmarking/getting-started, https://www.getmaxim.ai/bifrost/resources/benchmarks.md
- TensorZero: https://github.com/tensorzero/tensorzero, https://github.com/tensorzero/tensorzero/blob/main/docs/gateway/benchmarks.mdx
- Helicone AI Gateway: https://github.com/Helicone/ai-gateway, https://github.com/Helicone/ai-gateway/blob/main/benchmarks/README.md
- agentgateway: https://github.com/agentgateway/agentgateway, https://agentgateway.dev/blog/2026-06-26-benchmarking-agentgateway-vs-litellm-part-2/
- LiteLLM: https://docs.litellm.ai/docs/benchmarks, https://docs.litellm.ai/blog/rust-ai-gateway-benchmarks
- Portkey: https://github.com/Portkey-AI/gateway
- Kong: https://konghq.com/blog/engineering/ai-gateway-benchmark-kong-ai-gateway-portkey-litellm
- Envoy AI Gateway / Agent Router: https://github.com/envoyproxy/ai-gateway, https://theagentrouter.ai/docs/

Languages and runtimes:
- Go GC guide: https://tip.golang.org/doc/gc-guide
- Go 1.26 release notes: https://go.dev/doc/go1.26
- Java virtual threads: https://docs.oracle.com/en/java/javase/21/core/virtual-threads.html (JEP 491 at https://openjdk.org/jeps/491 could not be fetched)
- GraalVM native-image GC: https://www.graalvm.org/latest/reference-manual/native-image/optimizations-and-performance/MemoryManagement/
- .NET Native AOT: https://learn.microsoft.com/en-us/dotnet/core/deploying/native-aot/, https://learn.microsoft.com/en-us/dotnet/core/deploying/native-aot/cross-compile
- Bun executables: https://bun.com/docs/compiler/executables
- Node SEA: https://nodejs.org/api/single-executable-applications.html
- Zig: https://ziglang.org/download/
- Rust crates: https://crates.io/crates/age, https://crates.io/crates/webauthn-rs
- TechEmpower Round 23 (latest found, 2025-03-17; generic web-framework benchmark, not LLM proxying, so not used for numbers here): https://www.techempower.com/blog/2025/03/17/framework-benchmarks-round-23/

Unverified: Kong's Nginx/OpenResty base (not stated on the benchmark page); the CLIProxyAPI Rust port's maturity; rusqlite and Elixir library maturity; JEP 491 details.
