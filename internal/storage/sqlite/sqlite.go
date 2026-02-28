package sqlite

import (
	"database/sql"

	"github.com/Prohor722/go_rest_api_1/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	_,err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
}