package sql

import (
	_ "database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/golikoffegor/go-url-shortening-service/internal/interfaces"
	"github.com/golikoffegor/go-url-shortening-service/internal/model"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage/sql/postgresql"
	"github.com/golikoffegor/go-url-shortening-service/internal/utils"
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
	query := fmt.Sprintf("SELECT url, url_key, is_deleted FROM shortenerurls WHERE url_key = '%v';", key)
	rows := ps.db.RwDB.DB.QueryRow(query)
	err := rows.Scan(&shortening.URL, &shortening.Key, &shortening.IsDeleted)
	return &shortening, err
}

// Get возвращает URL из хранилища по ключу key
func (ps *PostgreSQLStorage) GetByURL(url string) (*model.Shortening, error) {
	shortening := model.Shortening{}
	query := fmt.Sprintf("SELECT url, url_key FROM shortenerurls WHERE url = '%v';", url)
	rows := ps.db.RwDB.DB.QueryRow(query)
	err := rows.Scan(&shortening.URL, &shortening.Key)
	return &shortening, err
}

// Get возвращает сохраненные URL пользователя из хранилища
func (ps *PostgreSQLStorage) GetByUserID(id string) ([]*model.Shortening, error) {
	shorteningList := []*model.Shortening{}
	query := fmt.Sprintf("SELECT url, url_key, user_id FROM shortenerurls WHERE user_id = '%v';", id)
	rows, err := ps.db.RwDB.DB.Query(query)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	for rows.Next() {
		shortening := model.Shortening{}
		_ = rows.Scan(&shortening.URL, &shortening.Key, &shortening.UserID)
		shorteningList = append(shorteningList, &shortening)
	}
	return shorteningList, err
}

// Delete удаляет сохраненные URL пользователя из хранилища по ключу
func (ps *PostgreSQLStorage) DeleteByUserIDBatch(doneCh chan struct{}, userID string, urlKeys []string) chan error {
	inputCh := utils.Generator(doneCh, urlKeys)

	return utils.FanIn(doneCh, utils.FanOut(10, func() chan error {
		result := make(chan error)

		go func() {
			defer close(result)

			for URL := range inputCh {
				query := fmt.Sprintf("UPDATE shortenerurls SET is_deleted = TRUE WHERE url_key = '%v' AND user_id = '%v';", URL, userID)
				_, err := ps.db.RwDB.DB.Exec(query)
				if err != nil {
					return
				}
				select {
				case <-doneCh:
					return
				case result <- err:
				}
			}
		}()

		return result
	}))
}

// Put записывает URL в хранилище с ключом key
func (ps *PostgreSQLStorage) Put(shortening model.Shortening) error {
	query := fmt.Sprintf("INSERT INTO shortenerurls (url, url_key, user_id) VALUES ('%v', '%v', '%v');", shortening.URL, shortening.Key, shortening.UserID)
	_, err := ps.db.RwDB.DB.Exec(query)
	if driverErr, ok := err.(*pgconn.PgError); ok && driverErr.Code == "23505" {
		return model.ConflictError{}
	}
	return err
}

func (ps *PostgreSQLStorage) PutBatch(shorteningList []model.Shortening) error {
	var values string
	for index, item := range shorteningList {
		values += fmt.Sprintf("('%v', '%v', '%v')", item.URL, item.Key, item.UserID)
		if index < len(shorteningList)-1 {
			values += ", "
		}
	}
	query := fmt.Sprintf("INSERT INTO shortenerurls (url, url_key, user_id) VALUES %s;", values)
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
