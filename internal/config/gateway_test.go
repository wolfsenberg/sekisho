package config

import (
	"testing"
	"time"
)

func TestGatewayPolicyOverridesAndBounds(t *testing.T) {
	env := map[string]string{
		"SEKISHO_GATEWAY_CAPACITY": "64", "SEKISHO_GATEWAY_QUEUE_LIMIT": "40",
		"SEKISHO_GATEWAY_QUEUE_TIMEOUT": "45s", "SEKISHO_GATEWAY_REQUEST_TIMEOUT": "20m",
		"SEKISHO_GATEWAY_STREAM_IDLE_TIMEOUT": "3m", "SEKISHO_GATEWAY_KEEPALIVE_INTERVAL": "5s",
		"SEKISHO_GATEWAY_RETRIES": "1", "SEKISHO_GATEWAY_RETRY_DELAY": "250ms",
		"SEKISHO_GATEWAY_TRANSIENT_THRESHOLD": "5", "SEKISHO_GATEWAY_TRANSIENT_WINDOW": "1m",
		"SEKISHO_GATEWAY_TRANSIENT_COOLDOWN": "15s", "SEKISHO_GATEWAY_RATE_LIMIT_COOLDOWN": "8s",
		"SEKISHO_INFERENCE_HEADER_TIMEOUT": "0s", "SEKISHO_SHUTDOWN_TIMEOUT": "45s",
	}
	if defaults, err := Load(func(string) string { return "" }); err != nil || defaults.ShutdownTimeout != 20*time.Second {
		t.Fatalf("default shutdown timeout: %v %v", defaults.ShutdownTimeout, err)
	}
	cfg, err := Load(func(k string) string { return env[k] })
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Gateway.Capacity != 64 || cfg.Gateway.QueueLimit != 40 || cfg.Gateway.QueueTimeout != 45*time.Second || cfg.Gateway.Retries != 1 || cfg.Gateway.RequestTimeout != 20*time.Minute || cfg.Gateway.StreamIdleTimeout != 3*time.Minute || cfg.InferenceHeaderTimeout != 0 || cfg.ShutdownTimeout != 45*time.Second {
		t.Fatalf("policy not applied: %+v", cfg)
	}
	for key, value := range map[string]string{
		"SEKISHO_GATEWAY_CAPACITY": "0", "SEKISHO_GATEWAY_QUEUE_LIMIT": "-1", "SEKISHO_GATEWAY_RETRIES": "20",
		"SEKISHO_GATEWAY_REQUEST_TIMEOUT": "0s", "SEKISHO_GATEWAY_QUEUE_TIMEOUT": "invalid", "SEKISHO_GATEWAY_TRANSIENT_WINDOW": "-1s",
		"SEKISHO_INFERENCE_HEADER_TIMEOUT": "-1s", "SEKISHO_SHUTDOWN_TIMEOUT": "11m",
	} {
		if _, err := Load(func(k string) string {
			if k == key {
				return value
			}
			return ""
		}); err == nil {
			t.Fatalf("invalid %s=%s accepted", key, value)
		}
	}
}
