package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/wolfsenberg/sekisho/internal/accounts"
	"github.com/wolfsenberg/sekisho/internal/gateway"
	"github.com/wolfsenberg/sekisho/internal/upstream"
)

// TestGatewayLoad measures the latency Sekisho itself adds to streamed Responses
// requests: API key authentication, pool and account selection, affinity, upstream
// dispatch, SSE relay and usage accounting, against an instant synthetic upstream.
// It is opt-in because it is a measurement, not a correctness check:
//
//	SEKISHO_LOAD=1 go test ./internal/server -run TestGatewayLoad -v
//
// Optional: SEKISHO_LOAD_SECONDS (per level, default 3), SEKISHO_LOAD_LEVELS
// (comma-separated concurrency, default 1,8,32,128), SEKISHO_LOAD_CHUNKS (SSE
// deltas per response, default 20), SEKISHO_LOAD_FAULT_PERCENT (share of upstream
// attempts answered with HTTP 500, default 0) and SEKISHO_LOAD_ACCOUNTS (pool size,
// default 1) to observe retries and failover under load.
func TestGatewayLoad(t *testing.T) {
	if os.Getenv("SEKISHO_LOAD") == "" {
		t.Skip("set SEKISHO_LOAD=1 to run the gateway load measurement")
	}
	seconds := envInt(t, "SEKISHO_LOAD_SECONDS", 3)
	chunks := envInt(t, "SEKISHO_LOAD_CHUNKS", 20)
	accountCount := envInt(t, "SEKISHO_LOAD_ACCOUNTS", 1)
	faultPercent := 0
	if os.Getenv("SEKISHO_LOAD_FAULT_PERCENT") != "" {
		faultPercent = envInt(t, "SEKISHO_LOAD_FAULT_PERCENT", 0)
	}
	levels := []int{1, 8, 32, 128}
	if raw := os.Getenv("SEKISHO_LOAD_LEVELS"); raw != "" {
		levels = nil
		for _, field := range strings.Split(raw, ",") {
			n, err := strconv.Atoi(strings.TrimSpace(field))
			if err != nil || n < 1 {
				t.Fatalf("invalid SEKISHO_LOAD_LEVELS entry %q", field)
			}
			levels = append(levels, n)
		}
	}

	var stream strings.Builder
	for i := range chunks {
		fmt.Fprintf(&stream, "data: {\"type\":\"response.output_text.delta\",\"item_id\":\"msg_load\",\"output_index\":0,\"content_index\":0,\"delta\":\"token %d \"}\n\n", i)
	}
	stream.WriteString("data: {\"type\":\"response.completed\",\"response\":{\"id\":\"resp_load\",\"output\":[],\"usage\":{\"input_tokens\":12,\"output_tokens\":24}}}\n\n")
	body := stream.String()

	policy := gateway.DefaultPolicy()
	policy.Capacity, policy.QueueLimit = 1000, 1000
	var attempts, faults atomic.Int64
	var injecting atomic.Bool
	fixture := newConfiguredForwardFixture(t, "codex", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		// Spread failures evenly (every k-th attempt) rather than in bursts.
		if n := attempts.Add(1); injecting.Load() && faultPercent > 0 && n%int64(max(1, 100/faultPercent)) == 0 {
			faults.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, body)
	}, false, policy)
	ctx := context.Background()
	for i := 1; i < accountCount; i++ {
		credential := accounts.Credential{AccessToken: "synthetic-upstream-access", RefreshToken: "synthetic-upstream-refresh", AccountID: fmt.Sprintf("synthetic-load-account-%d", i), ExpiresAt: time.Now().Add(time.Hour).Unix()}
		account, err := fixture.accounts.Authorize(ctx, fmt.Sprintf("Synthetic load %d", i), credential, "")
		if err != nil {
			t.Fatal(err)
		}
		if err := fixture.accounts.SaveCatalog(ctx, account.ID, 0, []string{"synthetic-model"}, time.Now().Unix(), upstream.CatalogSource("codex")); err != nil {
			t.Fatal(err)
		}
	}
	configureTestPool(t, fixture.connection)
	rows, err := fixture.accounts.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range rows {
		if err := fixture.forwarding.SetConcurrency(ctx, row.ID, 1000); err != nil {
			t.Fatal(err)
		}
	}
	// One synthetic member generates all load, so lift the default per-member limit
	// (10 concurrent) to measure the gateway rather than member admission.
	if err := fixture.forwarding.SetMemberLimits(ctx, fixture.userID, 0, 0); err != nil {
		t.Fatal(err)
	}

	client := &http.Client{Transport: &http.Transport{MaxIdleConns: 1024, MaxIdleConnsPerHost: 1024, MaxConnsPerHost: 0}}
	payload := `{"model":"synthetic-model","stream":true,"input":"Synthetic load input"}`
	call := func(session string) (time.Duration, time.Duration, error) {
		req, _ := http.NewRequest("POST", fixture.server.URL+"/v1/responses", strings.NewReader(payload))
		req.Header.Set("Authorization", "Bearer "+fixture.secret)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Session_id", session)
		start := time.Now()
		response, err := client.Do(req)
		if err != nil {
			return 0, 0, err
		}
		defer response.Body.Close()
		first := time.Since(start)
		data, err := io.ReadAll(response.Body)
		if err != nil {
			return 0, 0, err
		}
		if response.StatusCode != http.StatusOK || !strings.Contains(string(data), "response.completed") {
			return 0, 0, fmt.Errorf("status %d: %.200s", response.StatusCode, data)
		}
		return first, time.Since(start), nil
	}
	// Warm the catalog, affinity and connection paths before measuring.
	for i := range 3 {
		if _, _, err := call(fmt.Sprintf("load-warm-%d", i)); err != nil {
			t.Fatal("warmup:", err)
		}
	}
	injecting.Store(true)

	t.Logf("%-6s %9s %9s %9s %9s %9s %9s %7s %9s", "conc", "req/s", "ttfb p50", "p50", "p95", "p99", "max", "errors", "heap MB")
	for _, level := range levels {
		var (
			mu        sync.Mutex
			latencies []time.Duration
			firsts    []time.Duration
			failures  atomic.Int64
			firstErr  atomic.Value
			wg        sync.WaitGroup
		)
		deadline := time.Now().Add(time.Duration(seconds) * time.Second)
		started := time.Now()
		for worker := range level {
			wg.Add(1)
			go func() {
				defer wg.Done()
				local, localFirst := make([]time.Duration, 0, 1024), make([]time.Duration, 0, 1024)
				for n := 0; time.Now().Before(deadline); n++ {
					first, total, err := call(fmt.Sprintf("load-%d-%d-%d", level, worker, n%4))
					if err != nil {
						failures.Add(1)
						firstErr.CompareAndSwap(nil, err.Error())
						continue
					}
					local, localFirst = append(local, total), append(localFirst, first)
				}
				mu.Lock()
				latencies, firsts = append(latencies, local...), append(firsts, localFirst...)
				mu.Unlock()
			}()
		}
		wg.Wait()
		elapsed := time.Since(started)
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		slices.Sort(latencies)
		slices.Sort(firsts)
		t.Logf("%-6d %9.0f %9s %9s %9s %9s %9s %7d %9.1f", level, float64(len(latencies))/elapsed.Seconds(),
			pct(firsts, 50), pct(latencies, 50), pct(latencies, 95), pct(latencies, 99), pct(latencies, 100),
			failures.Load(), float64(mem.HeapInuse)/(1<<20))
		if msg := firstErr.Load(); msg != nil {
			t.Logf("       first error: %s", msg)
		}
	}
	if faultPercent > 0 {
		t.Logf("upstream attempts %d, injected failures %d, across %d accounts", attempts.Load(), faults.Load(), len(rows))
	}
}

func pct(sorted []time.Duration, p int) string {
	if len(sorted) == 0 {
		return "-"
	}
	i := (len(sorted)*p + 99) / 100
	if i > 0 {
		i--
	}
	return sorted[i].Round(10 * time.Microsecond).String()
}

func envInt(t *testing.T, name string, fallback int) int {
	t.Helper()
	raw := os.Getenv(name)
	if raw == "" {
		return fallback
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		t.Fatalf("invalid %s %q", name, raw)
	}
	return n
}
