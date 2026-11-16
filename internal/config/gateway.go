package config

import (
	"fmt"
	"strconv"
	"time"
)

func (c *Config) loadGateway(getenv func(string) string) error {
	p := &c.Gateway
	if value := getenv("SEKISHO_GATEWAY_QUOTA_WEIGHTED"); value != "" {
		if value != "true" && value != "false" {
			return fmt.Errorf("SEKISHO_GATEWAY_QUOTA_WEIGHTED must be true or false")
		}
		p.QuotaWeighted = value == "true"
	}
	for key, target := range map[string]*int{
		"SEKISHO_GATEWAY_CAPACITY": &p.Capacity, "SEKISHO_GATEWAY_QUEUE_LIMIT": &p.QueueLimit, "SEKISHO_GATEWAY_RETRIES": &p.Retries,
	} {
		if value := getenv(key); value != "" {
			n, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("%s must be an integer", key)
			}
			*target = n
		}
	}
	if value := getenv("SEKISHO_GATEWAY_TRANSIENT_THRESHOLD"); value != "" {
		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("SEKISHO_GATEWAY_TRANSIENT_THRESHOLD must be an integer")
		}
		p.TransientThreshold = n
	}
	for key, target := range map[string]*time.Duration{
		"SEKISHO_GATEWAY_QUEUE_TIMEOUT": &p.QueueTimeout, "SEKISHO_GATEWAY_REQUEST_TIMEOUT": &p.RequestTimeout,
		"SEKISHO_GATEWAY_STREAM_IDLE_TIMEOUT": &p.StreamIdleTimeout, "SEKISHO_GATEWAY_KEEPALIVE_INTERVAL": &p.KeepaliveInterval,
		"SEKISHO_GATEWAY_RETRY_DELAY": &p.RetryDelay, "SEKISHO_GATEWAY_TRANSIENT_WINDOW": &p.TransientWindow,
		"SEKISHO_GATEWAY_TRANSIENT_COOLDOWN": &p.TransientCooldown, "SEKISHO_GATEWAY_RATE_LIMIT_COOLDOWN": &p.RateLimitCooldown,
		"SEKISHO_INFERENCE_HEADER_TIMEOUT": &c.InferenceHeaderTimeout, "SEKISHO_SHUTDOWN_TIMEOUT": &c.ShutdownTimeout,
	} {
		if value := getenv(key); value != "" {
			d, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("%s must be a duration", key)
			}
			*target = d
		}
	}
	if c.InferenceHeaderTimeout < 0 || c.InferenceHeaderTimeout > 24*time.Hour {
		return fmt.Errorf("SEKISHO_INFERENCE_HEADER_TIMEOUT must be 0–24h")
	}
	if c.ShutdownTimeout < 0 || c.ShutdownTimeout > 10*time.Minute {
		return fmt.Errorf("SEKISHO_SHUTDOWN_TIMEOUT must be 0–10m")
	}
	return p.Validate()
}
