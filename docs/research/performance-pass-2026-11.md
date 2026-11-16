# Performance and reliability pass (November 2026)

Measured changes on top of upstream `v0.20.0` (`e361071`). Every number below comes from
`internal/server/load_test.go` on the same machine, so they compare to each other, not
to production hardware.

## How to reproduce

```sh
SEKISHO_LOAD=1 go test ./internal/server -run TestGatewayLoad -v
# optional: SEKISHO_LOAD_LEVELS=1,8,32,128 SEKISHO_LOAD_SECONDS=3 SEKISHO_LOAD_CHUNKS=20
# fault injection: SEKISHO_LOAD_FAULT_PERCENT=10 SEKISHO_LOAD_ACCOUNTS=3
```

The test drives the real HTTP server and the full request path: API key check, pool and
account selection, conversation affinity, upstream dispatch, SSE relay and usage
accounting. It uses a real SQLite file and an instant synthetic upstream that streams
20 deltas. The numbers are therefore Sekisho's own overhead; a real model adds hundreds
of milliseconds to minutes on top. The per-member concurrency limit (10 by default) is
lifted because one synthetic member generates all the load.

Environment: Linux (WSL2, kernel 6.6, ext4), 16 logical CPUs, Go 1.27.4.

## Results

| Concurrency | Before: req/s | Before: p50 | After: req/s | After: p50 |
|---|---|---|---|---|
| 1   | 45 | 19.6 ms | 273 | 3.2 ms |
| 8   | 47 | 159 ms  | 593 | 10.5 ms |
| 32  | 29 (51 rejected)  | 732 ms | 601 | 39 ms |
| 128 | 6 (213 rejected)  | 2.3 s  | 563 | 206 ms |

"Before" rejections were `member_busy`: that run kept the default member limit, so
requests piled up behind the slow gateway until the limit rejected them. Run-to-run
variation on this machine is roughly ±20%.

Time to first byte at concurrency 1 fell from 5.1 ms to 1.5 ms.

### Under upstream failures

With 10% of upstream attempts answered by HTTP 500 (`SEKISHO_LOAD_FAULT_PERCENT=10`),
the existing retry policy (2 retries, 500 ms delay) kept client-visible failures to
about 1 in 2,000–2,400 requests, with either 1 or 3 accounts in the pool. Retried
requests take about 500 ms longer (p95 ≈ 507 ms), which is
`SEKISHO_GATEWAY_RETRY_DELAY`.

## What changed and why

1. **SQLite `synchronous=NORMAL` in WAL mode** (`internal/storage/storage.go`).
   Profiling showed every request waiting for an fsync per commit, serialized on the
   single connection. NORMAL is SQLite's recommended setting for WAL: commits are
   safe from corruption and process crashes. The most recent commits can roll back
   after power loss or an OS crash.
2. **Credential writes stay durable against power loss** (`storage.Checkpoint`, called
   from `internal/accounts`). A rotated OAuth refresh token that rolled back would
   leave an account holding a token its provider already invalidated, so credential
   commits checkpoint the WAL to disk. These writes are rare (refresh about hourly
   per account, plus administrator imports).
3. **Prepared-statement cache per connection** (`internal/storage/statements.go`).
   sqlc sends every query as one-off SQL, so SQLite re-parsed and re-planned it each
   time. Parsing (`sqlite3Prepare`) was 28% of CPU, 3.5 times the cost of running
   the queries, and it happened inside the serialized section. The cache keeps one
   prepared statement per SQL text, at most 1,024. Generated code and call sites are
   unchanged. Backup verification opens the extracted copy with `storage.OpenSnapshot`,
   which skips the cache: `sqlite.Limit` needs the plain driver connection to bound
   values in untrusted archives.
4. **Graceful shutdown that actually drains** (`cmd/sekisho/main.go`). Request and
   gateway contexts derived from the signal context, so SIGTERM canceled every
   in-flight stream immediately. The `srv.Shutdown` grace period never applied. Now
   the signal only stops new connections. In-flight requests, including WebSocket
   sessions, get `SEKISHO_SHUTDOWN_TIMEOUT` (default 20s, below Compose's 30s
   `stop_grace_period`) before they are canceled. Cleanup then runs in the existing
   order.
5. **SQLite paths on Windows** (`storage.FileURI`). `file://C:/…` put the drive letter
   in the URI authority, so the server and most of the test suite could not open a
   database on Windows. This does not affect the Linux release images.
6. **Larger upstream idle connection pool** (`internal/upstream/client.go`), from
   16 total / 8 per host to 256 / 64. HTTP/1.1 API channels with more than 8
   concurrent streams were opening a fresh TLS connection per extra request.

## Verification

- `go vet ./...` is clean. `go test ./...` passes on Linux: 27 of 27 packages.
- Not run: the `-race` suite, because no C compiler was available in the test
  environment. Run `make test` on CI or a Linux machine with gcc before release.
- On Windows, tests that depend on POSIX file modes (`0600` key checks), on symlinks,
  or on `/bin/sh` (one frontend client test) fail for environmental reasons. The
  release targets are Linux images.
- New tests: graceful drain and grace-period cancellation (`cmd/sekisho/serve_test.go`),
  `SEKISHO_SHUTDOWN_TIMEOUT` bounds, WAL `NORMAL` plus checkpoint durability, the
  statement cache (reuse, the same SQL nested inside its own open rows, schema changes,
  transactions), and Windows file URIs.

## Deliberately not changed

- **API key, membership and pool checks are not cached.** `AGENTS.md` requires every
  request and WebSocket turn to recheck them.
- **The gateway mutex still covers the usage charge in `observation.finish`.** The
  code says admission must not see a released lease before the token charge commits.
  Releasing it earlier could let a member overspend an allowance. With the two SQLite
  fixes, the time spent under that lock is now small.
- **Usage writes are not batched.** History and aggregates must commit together with
  the charge, and batching would delay or weaken that.

## Remaining bottleneck

Throughput levels off at about 600 req/s, because all database work still goes
through one SQLite connection and the gateway mutex. That is well above what a team
gateway sees: 100 developers running agents rarely exceed tens of requests per second.
The next step, if ever needed, would be a separate read-only connection pool for
queries outside transactions. That needs care with the "no request may hold a
transaction during upstream IO" rule and with read-after-write expectations, so it
should be done only once a real workload shows the need.
