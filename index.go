package data_sqlite

import (
	"github.com/bamgoo/bamgoo"
	"github.com/bamgoo/data"
)

func Driver() data.Driver {
	return &sqliteDriver{}
}

func init() {
	bamgoo.Register("sqlite", Driver())
	bamgoo.Register("sqlite3", Driver())
}
