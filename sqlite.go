package data_sqlite

import (
	"database/sql"
	"strings"
	"sync/atomic"

	"github.com/infrago/data"
	_ "modernc.org/sqlite"
)

type (
	sqliteDriver struct{}

	sqliteConnection struct {
		instance *data.Instance
		db       *sql.DB
		actives  int64
	}

	sqliteDialect struct{}
)

func (d *sqliteDriver) Connect(inst *data.Instance) (data.Connection, error) {
	return &sqliteConnection{instance: inst}, nil
}

func (c *sqliteConnection) Open() error {
	dsn := strings.TrimSpace(c.instance.Config.Url)
	if dsn == "" {
		if v, ok := c.instance.Setting["dsn"].(string); ok {
			dsn = v
		}
	}
	if dsn == "" {
		dsn = "file:data.db"
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return err
	}
	c.db = db
	return nil
}

func (c *sqliteConnection) Close() error {
	if c.db == nil {
		return nil
	}
	err := c.db.Close()
	c.db = nil
	return err
}

func (c *sqliteConnection) Health() data.Health {
	return data.Health{Workload: atomic.LoadInt64(&c.actives)}
}

func (c *sqliteConnection) DB() *sql.DB {
	return c.db
}

func (c *sqliteConnection) Dialect() data.Dialect {
	return sqliteDialect{}
}

func (sqliteDialect) Name() string { return "sqlite" }
func (sqliteDialect) Quote(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, `"`, ``)
	return `"` + s + `"`
}
func (sqliteDialect) Placeholder(_ int) string { return "?" }
func (sqliteDialect) SupportsILike() bool      { return false }
func (sqliteDialect) SupportsReturning() bool  { return false }
