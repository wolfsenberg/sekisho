package storage

import (
	"context"
	"path/filepath"
	"testing"
)

func TestDatabaseUsesWALNormalSyncAndCheckpointFlushesCommits(t *testing.T) {
	ctx := context.Background()
	db, err := Open(ctx, filepath.Join(t.TempDir(), "sekisho.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var mode string
	var synchronous int
	if err := db.QueryRowContext(ctx, "PRAGMA journal_mode").Scan(&mode); err != nil || mode != "wal" {
		t.Fatalf("journal mode %q %v", mode, err)
	}
	// NORMAL (1): WAL commits survive process crashes without an fsync per commit.
	if err := db.QueryRowContext(ctx, "PRAGMA synchronous").Scan(&synchronous); err != nil || synchronous != 1 {
		t.Fatalf("synchronous %d %v", synchronous, err)
	}
	if _, err := db.ExecContext(ctx, "CREATE TABLE durability_probe(value TEXT)"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "INSERT INTO durability_probe VALUES('rotated credential')"); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(ctx, db); err != nil {
		t.Fatal(err)
	}
	var busy, logged, copied int
	if err := db.QueryRowContext(ctx, "PRAGMA wal_checkpoint(PASSIVE)").Scan(&busy, &logged, &copied); err != nil {
		t.Fatal(err)
	}
	if busy != 0 || logged != copied {
		t.Fatalf("commits remain only in the WAL after Checkpoint: busy=%d logged=%d copied=%d", busy, logged, copied)
	}
}
