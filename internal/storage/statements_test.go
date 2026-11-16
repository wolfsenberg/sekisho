package storage

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"
)

func openCached(t *testing.T) (*sql.DB, *cachingConn) {
	t.Helper()
	db, err := Open(context.Background(), filepath.Join(t.TempDir(), "sekisho.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var cached *cachingConn
	if err := conn.Raw(func(raw any) error {
		cached, _ = raw.(*cachingConn)
		return nil
	}); err != nil || cached == nil {
		t.Fatalf("storage.Open connection is not statement-caching: %T %v", cached, err)
	}
	conn.Close()
	return db, cached
}

func TestRepeatedQueriesReuseOnePreparedStatement(t *testing.T) {
	ctx := context.Background()
	db, cached := openCached(t)
	if _, err := db.ExecContext(ctx, "CREATE TABLE probe(id INTEGER PRIMARY KEY, value TEXT)"); err != nil {
		t.Fatal(err)
	}
	for i := range 3 {
		if _, err := db.ExecContext(ctx, "INSERT INTO probe(value) VALUES(?)", "row"); err != nil {
			t.Fatal(err)
		}
		var n int
		if err := db.QueryRowContext(ctx, "SELECT count(*) FROM probe WHERE value = ?", "row").Scan(&n); err != nil || n != i+1 {
			t.Fatalf("count %d %v", n, err)
		}
	}
	cached.mu.Lock()
	defer cached.mu.Unlock()
	for _, query := range []string{"INSERT INTO probe(value) VALUES(?)", "SELECT count(*) FROM probe WHERE value = ?"} {
		if cached.stmts[query] == nil {
			t.Fatalf("statement was not cached: %q", query)
		}
	}
}

func TestCachedStatementNestedInItsOwnOpenRowsStaysCorrect(t *testing.T) {
	ctx := context.Background()
	db, _ := openCached(t)
	if _, err := db.ExecContext(ctx, "CREATE TABLE probe(id INTEGER PRIMARY KEY, parent INTEGER)"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "INSERT INTO probe(id, parent) VALUES (1, 0), (2, 1), (3, 1), (4, 2)"); err != nil {
		t.Fatal(err)
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	const children = "SELECT id FROM probe WHERE parent = ? ORDER BY id"
	outer, err := tx.QueryContext(ctx, children, 1)
	if err != nil {
		t.Fatal(err)
	}
	defer outer.Close()
	var got []int
	for outer.Next() {
		var id int
		if err := outer.Scan(&id); err != nil {
			t.Fatal(err)
		}
		// The same SQL runs while the outer rows still hold its cached statement.
		inner, err := tx.QueryContext(ctx, children, id)
		if err != nil {
			t.Fatal(err)
		}
		for inner.Next() {
			var child int
			if err := inner.Scan(&child); err != nil {
				t.Fatal(err)
			}
			got = append(got, child)
		}
		if err := inner.Close(); err != nil {
			t.Fatal(err)
		}
		got = append(got, id)
	}
	if err := outer.Err(); err != nil {
		t.Fatal(err)
	}
	if want := []int{4, 2, 3}; len(got) != len(want) || got[0] != want[0] || got[1] != want[1] || got[2] != want[2] {
		t.Fatalf("nested results %v, want %v", got, want)
	}
}

func TestCachedStatementsFollowSchemaChanges(t *testing.T) {
	ctx := context.Background()
	db, _ := openCached(t)
	if _, err := db.ExecContext(ctx, "CREATE TABLE probe(a INTEGER)"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "INSERT INTO probe(a) VALUES (1)"); err != nil {
		t.Fatal(err)
	}
	columns := func() []string {
		rows, err := db.QueryContext(ctx, "SELECT * FROM probe")
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()
		names, err := rows.Columns()
		if err != nil {
			t.Fatal(err)
		}
		return names
	}
	if got := columns(); len(got) != 1 {
		t.Fatalf("columns before migration %v", got)
	}
	if _, err := db.ExecContext(ctx, "ALTER TABLE probe ADD COLUMN b TEXT NOT NULL DEFAULT 'x'"); err != nil {
		t.Fatal(err)
	}
	if got := columns(); len(got) != 2 {
		t.Fatalf("cached statement kept the old schema: %v", got)
	}
}

func TestCachedStatementsWorkInsideTransactions(t *testing.T) {
	ctx := context.Background()
	db, _ := openCached(t)
	if _, err := db.ExecContext(ctx, "CREATE TABLE probe(value INTEGER)"); err != nil {
		t.Fatal(err)
	}
	for _, commit := range []bool{false, true} {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := tx.ExecContext(ctx, "INSERT INTO probe(value) VALUES (?)", 7); err != nil {
			t.Fatal(err)
		}
		if commit {
			err = tx.Commit()
		} else {
			err = tx.Rollback()
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	var n int
	if err := db.QueryRowContext(ctx, "SELECT count(*) FROM probe WHERE value = ?", 7).Scan(&n); err != nil || n != 1 {
		t.Fatalf("transaction results %d %v", n, err)
	}
}
