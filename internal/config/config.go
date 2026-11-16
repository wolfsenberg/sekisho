package config

import (
	"fmt"
	"log/slog"
	"math"
	"net"
	"net/netip"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/wolfsenberg/sekisho/internal/backup"
	"github.com/wolfsenberg/sekisho/internal/gateway"
	"github.com/wolfsenberg/sekisho/internal/vault"
)

type Config struct {
	Backup                 backup.Schedule
	BackupWebhookURL       string
	Demo                   bool
	Addr                   string
	DataDir                string
	CredentialsKeyFile     string
	LogLevel               slog.Level
	MaxRequestBody         int64
	PublicURL              string
	TrustedProxies         []netip.Prefix
	PricingURL             string
	PricingHashURL         string
	PricingCache           string
	PricingFallback        string
	PricingOverride        string
	PricingInterval        time.Duration
	MetricsToken           string
	UpdateDir              string
	Gateway                gateway.Policy
	InferenceHeaderTimeout time.Duration
	// ShutdownTimeout bounds how long in-flight requests may finish after a stop signal.
	ShutdownTimeout time.Duration
}

func Load(getenv func(string) string) (Config, error) {
	c := Config{Addr: "127.0.0.1:8080", DataDir: "./data", LogLevel: slog.LevelInfo, MaxRequestBody: 128 << 20, PricingInterval: 6 * time.Hour, Gateway: gateway.DefaultPolicy(), InferenceHeaderTimeout: 5 * time.Minute, ShutdownTimeout: 20 * time.Second}
	c.UpdateDir = getenv("SEKISHO_UPDATE_DIR")
	if c.UpdateDir != "" && !filepath.IsAbs(c.UpdateDir) {
		return c, fmt.Errorf("SEKISHO_UPDATE_DIR must be an absolute host-worker mailbox path")
	}
	if err := c.loadGateway(getenv); err != nil {
		return c, err
	}
	if err := c.loadBackup(getenv); err != nil {
		return c, err
	}
	if v := getenv("SEKISHO_DEMO"); v != "" {
		if v != "true" && v != "false" {
			return c, fmt.Errorf("SEKISHO_DEMO must be true or false")
		}
		c.Demo = v == "true"
	}
	if v := getenv("SEKISHO_ADDR"); v != "" {
		c.Addr = v
	}
	if v := getenv("SEKISHO_DATA_DIR"); v != "" {
		c.DataDir = v
	}
	if v := getenv("SEKISHO_CREDENTIALS_KEY_FILE"); v != "" {
		if !c.Demo {
			if _, err := vault.KeyPath(c.DataDir, v); err != nil {
				return c, fmt.Errorf("SEKISHO_CREDENTIALS_KEY_FILE: %w", err)
			}
		}
		c.CredentialsKeyFile = v
	}
	if v := getenv("SEKISHO_LOG_LEVEL"); v != "" {
		if err := c.LogLevel.UnmarshalText([]byte(v)); err != nil {
			return c, fmt.Errorf("SEKISHO_LOG_LEVEL: %w", err)
		}
	}
	if v := getenv("SEKISHO_MAX_REQUEST_BODY_MB"); v != "" {
		mb, err := strconv.ParseInt(v, 10, 64)
		if err != nil || mb <= 0 || mb > math.MaxInt64>>20 {
			return c, fmt.Errorf("SEKISHO_MAX_REQUEST_BODY_MB must be a positive integer fitting in bytes")
		}
		c.MaxRequestBody = mb << 20
	}
	if v := getenv("SEKISHO_PUBLIC_URL"); v != "" {
		u, err := url.Parse(v)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" || u.User != nil || (u.Path != "" && u.Path != "/") || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" {
			return c, fmt.Errorf("SEKISHO_PUBLIC_URL must be an http(s) origin without credentials, path, query, or fragment")
		}
		host := strings.ToLower(u.Host)
		if port := u.Port(); port != "" {
			n, err := strconv.Atoi(port)
			if err != nil || n < 1 || n > 65535 {
				return c, fmt.Errorf("invalid SEKISHO_PUBLIC_URL port")
			}
			if (u.Scheme == "https" && port == "443") || (u.Scheme == "http" && port == "80") {
				host = strings.ToLower(u.Hostname())
				if strings.Contains(host, ":") {
					host = "[" + host + "]"
				}
			}
		}
		c.PublicURL = u.Scheme + "://" + host
	}
	if v := getenv("SEKISHO_TRUSTED_PROXIES"); v != "" {
		for _, part := range strings.Split(v, ",") {
			prefix, err := netip.ParsePrefix(strings.TrimSpace(part))
			if err != nil {
				return c, fmt.Errorf("SEKISHO_TRUSTED_PROXIES must be comma-separated IP CIDRs: %w", err)
			}
			c.TrustedProxies = append(c.TrustedProxies, prefix.Masked())
		}
	}
	c.PricingURL = getenv("SEKISHO_PRICING_URL")
	c.MetricsToken = getenv("SEKISHO_METRICS_TOKEN")
	if c.MetricsToken != "" {
		if len(c.MetricsToken) < 32 || len(c.MetricsToken) > 128 {
			return c, fmt.Errorf("SEKISHO_METRICS_TOKEN must contain 32–128 visible ASCII characters")
		}
		for _, char := range c.MetricsToken {
			if char < 33 || char > 126 {
				return c, fmt.Errorf("SEKISHO_METRICS_TOKEN must contain 32–128 visible ASCII characters")
			}
		}
	}
	c.PricingHashURL = getenv("SEKISHO_PRICING_HASH_URL")
	c.PricingCache = getenv("SEKISHO_PRICING_CACHE")
	c.PricingFallback = getenv("SEKISHO_PRICING_FALLBACK")
	c.PricingOverride = getenv("SEKISHO_PRICING_OVERRIDE")
	if v := getenv("SEKISHO_PRICING_INTERVAL"); v != "" {
		interval, err := time.ParseDuration(v)
		if err != nil || interval <= 0 {
			return c, fmt.Errorf("SEKISHO_PRICING_INTERVAL must be a positive duration")
		}
		c.PricingInterval = interval
	}
	_, port, err := net.SplitHostPort(c.Addr)
	if err != nil {
		return c, fmt.Errorf("SEKISHO_ADDR must be host:port: %w", err)
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return c, fmt.Errorf("SEKISHO_ADDR port must be between 1 and 65535")
	}
	return c, nil
}
