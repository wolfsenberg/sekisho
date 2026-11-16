package backup

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/wolfsenberg/sekisho/internal/storage"
	"github.com/wolfsenberg/sekisho/internal/turnstile"
	"github.com/wolfsenberg/sekisho/internal/vault"
)

func TestBackupVerifiesTurnstileConfiguration(t *testing.T) {
	ctx := context.Background()
	dir := t.TempDir()
	conn, err := storage.Open(ctx, filepath.Join(dir, databaseName))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	path := filepath.Join(dir, keyName)
	v, _ := vault.Open(path, true)
	s, err := turnstile.New(ctx, conn, v, nil)
	if err != nil {
		t.Fatal(err)
	}
	secret := "synthetic-secret"
	if _, err := s.Save(ctx, turnstile.Input{Enabled: true, SiteKey: "synthetic-site", SecretKey: &secret}); err != nil {
		t.Fatal(err)
	}
	if err := verifyCredentials(ctx, conn, path); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Exec("UPDATE settings SET value='corrupt-synthetic-config' WHERE key='instance.turnstile'"); err != nil {
		t.Fatal(err)
	}
	if err := verifyCredentials(ctx, conn, path); err == nil {
		t.Fatal("accepted undecryptable Turnstile configuration")
	}
}
