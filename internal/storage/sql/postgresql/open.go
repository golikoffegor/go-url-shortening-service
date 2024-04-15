package postgresql

import (
	// postgresql database driver
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/golikoffegor/go-url-shortening-service/config"
)

type DBConnection struct {
	ConnString string
	RwDB       *sqlx.DB
	RoDB       *sqlx.DB
}

func Open() (*DBConnection, error) {
	db, err := sqlx.Open("pgx", config.PostgreSQLDSN)
	if err != nil {
		return nil, err
	}

	return &DBConnection{
		ConnString: config.PostgreSQLDSN,
		RwDB:       db,
		RoDB:       db,
	}, nil
}
