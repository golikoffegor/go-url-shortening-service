package storage

import (
	"github.com/golikoffegor/go-url-shortening-service/config"
	"github.com/golikoffegor/go-url-shortening-service/internal/interfaces"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage/filestorage"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage/memstorage"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage/sql"
)

func GetStorage() interfaces.Storager {
	if config.PostgreSQLDSN != "" {
		return sql.NewStorage()
	}
	if config.FileStoragePath != "" {
		return filestorage.NewStorage()
	}
	return memstorage.NewStorage()
}
