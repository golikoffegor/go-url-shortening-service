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

// NewStorage создает новый экземпляр MemStorage
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

func (m MemStorage) PutBatch(shorteningList []model.Shortening) error {
	for _, item := range shorteningList {
		m.urlAddresses[item.Key] = item.URL
	}
	return nil
}

// Initialize хранилища
func (m MemStorage) Initialize() error {
	return nil
}

// TODO
func (m MemStorage) GetByURL(url string) (*model.Shortening, error) {
	return nil, nil
}

// TODO
func (m MemStorage) GetByUserID(id string) ([]*model.Shortening, error) {
	return nil, nil
}

// TODO
func (m MemStorage) DeleteByUserIDBatch(doneCh chan struct{}, userID string, urlKeys []string) chan error {
	return nil
}
