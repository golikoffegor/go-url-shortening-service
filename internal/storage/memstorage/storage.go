package memstorage

import (
	"errors"

	"github.com/golikoffegor/go-url-shortening-service/internal/interfaces"
	"github.com/golikoffegor/go-url-shortening-service/internal/model"
)

// MemStorage хранилище
type MemStorage struct {
	urlAddresses map[string]string
}

// NewMemStorage создает новый экземпляр MemStorage
func NewStorage() interfaces.Storager {
	return &MemStorage{
		urlAddresses: make(map[string]string),
	}
}

// Get возвращает URL из хранилища по ключу key
func (m MemStorage) Get(key string) (*model.Shortening, error) {
	v, ok := m.urlAddresses[key]
	if !ok {
		return nil, errors.New("url not found")
	}
	shortening := model.Shortening{Key: key, URL: v}
	return &shortening, nil
}

// Put записывает URL в хранилище с ключом key
func (m MemStorage) Put(shortening model.Shortening) error {
	m.urlAddresses[shortening.Key] = shortening.URL
	return nil
}
