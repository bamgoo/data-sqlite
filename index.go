package data_sqlite

import (
	"github.com/infrago/infra"
	"github.com/infrago/data"
)

func Driver() data.Driver {
	return &sqliteDriver{}
}

func init() {
	infra.Register("sqlite", Driver())
	infra.Register("sqlite3", Driver())
}
