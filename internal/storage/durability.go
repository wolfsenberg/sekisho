package storage

import (
	"context"
	"database/sql"
	"fmt"
)

// Checkpoint copies committed WAL frames into the main database file, syncing both.
//
// The database runs in WAL mode with synchronous=NORMAL: commits survive process
// crashes without an fsync each, but the most recent ones may roll back after power
// loss or an OS crash until the next checkpoint. Call Checkpoint after commits that
// must survive that, such as rotated OAuth credentials whose predecessors the provider
// has already invalidated.
func Checkpoint(ctx context.Context, db *sql.DB) error {
	var busy, logged, copied int
	if err := db.QueryRowContext(ctx, "PRAGMA wal_checkpoint(PASSIVE)").Scan(&busy, &logged, &copied); err != nil {
		return err
	}
	if busy != 0 || logged != copied {
		return fmt.Errorf("checkpoint incomplete: %d of %d WAL frames copied", copied, logged)
	}
	return nil
}
