package sql

import (
	_ "database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/golikoffegor/go-url-shortening-service/internal/interfaces"
	"github.com/golikoffegor/go-url-shortening-service/internal/model"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage/sql/postgresql"
)

// PostgreSQLStorage хранилище
type PostgreSQLStorage struct {
	db *postgresql.DBConnection
}

// NewStorage создает новый экземпляр PostgreSQLStorage
func NewStorage() interfaces.Storager {
	return &PostgreSQLStorage{}
}

// Get возвращает URL из хранилища по ключу key
func (ps *PostgreSQLStorage) Get(key string) (*model.Shortening, error) {
	shortening := model.Shortening{}
	query := fmt.Sprintf("SELECT url, url_key FROM shortenerurls WHERE url_key = '%v';", key)
	rows := ps.db.RwDB.DB.QueryRow(query)
	err := rows.Scan(&shortening.URL, &shortening.Key)
	return &shortening, err
}

// Put записывает URL в хранилище с ключом key
func (ps *PostgreSQLStorage) Put(shortening model.Shortening) error {
	query := fmt.Sprintf("INSERT INTO shortenerurls (url, url_key) VALUES ('%v', '%v');", shortening.URL, shortening.Key)
	_, err := ps.db.RwDB.DB.Exec(query)
	return err
}

func (ps *PostgreSQLStorage) PutBatch(shorteningList []model.Shortening) error {
	var values string
	for index, item := range shorteningList {
		values += fmt.Sprintf("('%v', '%v')", item.URL, item.Key)
		if index < len(shorteningList)-1 {
			values += ", "
		}
	}
	query := fmt.Sprintf("INSERT INTO shortenerurls (url, url_key) VALUES %s;", values)
	_, err := ps.db.RwDB.DB.Exec(query)
	return err
}

func MigrateUp(driverName string, db *sqlx.DB) (int, error) {
	migrations := migrate.AssetMigrationSource{
		Asset:    postgresql.Asset,
		AssetDir: postgresql.AssetDir,
		Dir:      "internal/storage/sql/postgresql/migrations/default",
	}

	return migrate.Exec(db.DB, driverName, migrations, migrate.Up)
}

// Initialize хранилища
func (ps *PostgreSQLStorage) Initialize() error {
	db, err := postgresql.Open()
	if err != nil {
		return err
	}
	ps.db = db
	_, err = MigrateUp("postgres", db.RwDB)
	return err
}
