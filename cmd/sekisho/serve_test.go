package main

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"testing"
	"time"
)

type serveRun struct {
	url     string
	stop    context.CancelFunc
	runtime context.Context
	done    chan error
}

func startServe(t *testing.T, handler http.Handler, grace time.Duration) serveRun {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	stopping, stop := context.WithCancel(context.Background())
	runtime, cancelRuntime := context.WithCancel(context.Background())
	t.Cleanup(func() { stop(); cancelRuntime() })
	done := make(chan error, 1)
	go func() {
		done <- serve(stopping, runtime, cancelRuntime, slog.New(slog.NewTextHandler(io.Discard, nil)), listener, handler, grace)
	}()
	return serveRun{url: "http://" + listener.Addr().String(), stop: stop, runtime: runtime, done: done}
}

func TestServeDrainsInFlightRequestsBeforeCancellingRuntime(t *testing.T) {
	started, release := make(chan struct{}), make(chan struct{})
	canceled := make(chan bool, 1)
	run := startServe(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		select {
		case <-release:
			canceled <- false
			io.WriteString(w, "complete")
		case <-r.Context().Done():
			canceled <- true
		}
	}), 5*time.Second)
	response := make(chan string, 1)
	go func() {
		res, err := http.Get(run.url)
		if err != nil {
			response <- "error: " + err.Error()
			return
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		response <- string(body)
	}()
	<-started
	run.stop()
	// A stop signal must not cancel work that is already streaming.
	time.Sleep(150 * time.Millisecond)
	if run.runtime.Err() != nil {
		t.Fatal("runtime context was canceled before in-flight requests drained")
	}
	close(release)
	if wasCanceled := <-canceled; wasCanceled {
		t.Fatal("in-flight request was canceled by the stop signal")
	}
	if body := <-response; body != "complete" {
		t.Fatalf("in-flight response = %q", body)
	}
	select {
	case err := <-run.done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("serve did not return after draining")
	}
	if run.runtime.Err() == nil {
		t.Fatal("runtime context must be canceled once serving stops")
	}
}

func TestServeCancelsRequestsThatOutliveShutdownGrace(t *testing.T) {
	started := make(chan struct{})
	run := startServe(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		<-r.Context().Done()
	}), 100*time.Millisecond)
	go func() {
		if res, err := http.Get(run.url); err == nil {
			res.Body.Close()
		}
	}()
	<-started
	begin := time.Now()
	run.stop()
	select {
	case <-run.done:
	case <-time.After(5 * time.Second):
		t.Fatal("serve did not stop a request that outlived the grace period")
	}
	if elapsed := time.Since(begin); elapsed < 100*time.Millisecond {
		t.Fatalf("serve canceled requests after %s, before the grace period", elapsed)
	}
	if run.runtime.Err() == nil {
		t.Fatal("runtime context was not canceled after the grace period")
	}
}
