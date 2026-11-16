package storage

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"
	"sync"
)

// maxCachedStatements bounds the per-connection cache. Application queries are a
// fixed set of sqlc statements, so the bound only matters for ad hoc SQL.
const maxCachedStatements = 1024

// openCaching opens SQLite through a connector whose connections keep one prepared
// statement per SQL text. sqlc issues every query as one-off SQL, and SQLite then
// re-parses and re-plans it each time; with all access serialized on one connection,
// that parsing was the largest share of gateway CPU under load.
func openCaching(dsn string) (*sql.DB, error) {
	base, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	sqlite := base.Driver()
	if err := base.Close(); err != nil {
		return nil, err
	}
	return sql.OpenDB(cachingConnector{dsn: dsn, driver: sqlite}), nil
}

type cachingConnector struct {
	dsn    string
	driver driver.Driver
}

func (c cachingConnector) Connect(context.Context) (driver.Conn, error) {
	conn, err := c.driver.Open(c.dsn)
	if err != nil {
		return nil, err
	}
	return &cachingConn{Conn: conn, stmts: make(map[string]*cachedStmt)}, nil
}

func (c cachingConnector) Driver() driver.Driver { return c.driver }

type cachedStmt struct {
	stmt driver.Stmt
	busy bool // its rows are still open; reusing it now would reset them
}

type cachingConn struct {
	driver.Conn
	mu    sync.Mutex
	stmts map[string]*cachedStmt
}

// acquire returns a prepared statement for query and a release function. A statement
// whose rows are still open (the same SQL nested inside its own result iteration) is
// never shared; that caller gets a temporary statement instead.
func (c *cachingConn) acquire(ctx context.Context, query string) (driver.Stmt, func(), error) {
	c.mu.Lock()
	entry := c.stmts[query]
	if entry != nil && !entry.busy {
		entry.busy = true
		c.mu.Unlock()
		return entry.stmt, func() {
			c.mu.Lock()
			entry.busy = false
			c.mu.Unlock()
		}, nil
	}
	cache := entry == nil && len(c.stmts) < maxCachedStatements
	c.mu.Unlock()
	stmt, err := c.PrepareContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	if !cache {
		return stmt, func() { _ = stmt.Close() }, nil
	}
	entry = &cachedStmt{stmt: stmt, busy: true}
	c.mu.Lock()
	c.stmts[query] = entry
	c.mu.Unlock()
	return stmt, func() {
		c.mu.Lock()
		entry.busy = false
		c.mu.Unlock()
	}, nil
}

func (c *cachingConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	stmt, release, err := c.acquire(ctx, query)
	if err != nil {
		return nil, err
	}
	defer release()
	exec, ok := stmt.(driver.StmtExecContext)
	if !ok {
		return nil, driver.ErrSkip
	}
	return exec.ExecContext(ctx, args)
}

func (c *cachingConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	stmt, release, err := c.acquire(ctx, query)
	if err != nil {
		return nil, err
	}
	queryer, ok := stmt.(driver.StmtQueryContext)
	if !ok {
		release()
		return nil, driver.ErrSkip
	}
	rows, err := queryer.QueryContext(ctx, args)
	if err != nil {
		release()
		return nil, err
	}
	return &releasingRows{Rows: rows, release: release}, nil
}

func (c *cachingConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if p, ok := c.Conn.(driver.ConnPrepareContext); ok {
		return p.PrepareContext(ctx, query)
	}
	return c.Conn.Prepare(query)
}

func (c *cachingConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if b, ok := c.Conn.(driver.ConnBeginTx); ok {
		return b.BeginTx(ctx, opts)
	}
	return c.Conn.Begin()
}

func (c *cachingConn) Ping(ctx context.Context) error {
	if p, ok := c.Conn.(driver.Pinger); ok {
		return p.Ping(ctx)
	}
	return nil
}

func (c *cachingConn) ResetSession(ctx context.Context) error {
	if r, ok := c.Conn.(driver.SessionResetter); ok {
		return r.ResetSession(ctx)
	}
	return nil
}

func (c *cachingConn) IsValid() bool {
	if v, ok := c.Conn.(driver.Validator); ok {
		return v.IsValid()
	}
	return true
}

func (c *cachingConn) Close() error {
	c.mu.Lock()
	var errs []error
	for query, entry := range c.stmts {
		errs = append(errs, entry.stmt.Close())
		delete(c.stmts, query)
	}
	c.mu.Unlock()
	return errors.Join(append(errs, c.Conn.Close())...)
}

// releasingRows returns its cached statement on Close and forwards the optional
// column metadata interfaces of the underlying rows.
type releasingRows struct {
	driver.Rows
	release func()
	once    sync.Once
}

func (r *releasingRows) Close() error {
	err := r.Rows.Close()
	r.once.Do(r.release)
	return err
}

func (r *releasingRows) ColumnTypeDatabaseTypeName(index int) string {
	if c, ok := r.Rows.(driver.RowsColumnTypeDatabaseTypeName); ok {
		return c.ColumnTypeDatabaseTypeName(index)
	}
	return ""
}

func (r *releasingRows) ColumnTypeLength(index int) (int64, bool) {
	if c, ok := r.Rows.(driver.RowsColumnTypeLength); ok {
		return c.ColumnTypeLength(index)
	}
	return 0, false
}

func (r *releasingRows) ColumnTypeNullable(index int) (bool, bool) {
	if c, ok := r.Rows.(driver.RowsColumnTypeNullable); ok {
		return c.ColumnTypeNullable(index)
	}
	return false, false
}

func (r *releasingRows) ColumnTypePrecisionScale(index int) (int64, int64, bool) {
	if c, ok := r.Rows.(driver.RowsColumnTypePrecisionScale); ok {
		return c.ColumnTypePrecisionScale(index)
	}
	return 0, 0, false
}

func (r *releasingRows) ColumnTypeScanType(index int) reflect.Type {
	if c, ok := r.Rows.(driver.RowsColumnTypeScanType); ok {
		return c.ColumnTypeScanType(index)
	}
	return reflect.TypeFor[any]()
}
