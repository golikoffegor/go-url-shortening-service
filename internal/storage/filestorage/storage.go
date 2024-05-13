package filestorage

import (
	"errors"
	"sync"

	"github.com/golikoffegor/go-url-shortening-service/config"
	"github.com/golikoffegor/go-url-shortening-service/internal/interfaces"
	"github.com/golikoffegor/go-url-shortening-service/internal/model"
)

type FileStorage struct {
	m sync.Map
}

func NewStorage() interfaces.Storager {
	return &FileStorage{}
}

func (fs *FileStorage) Put(shortening model.Shortening) error {
	if _, exists := fs.m.Load(shortening.Key); exists {
		return errors.New("identifier already exists")
	}

	Writer, err := NewWriter(config.FileStoragePath)
	if err != nil {
		return err
	}
	defer Writer.Close()

	if err := Writer.WriteFile(&model.File{
		ShortURL:    shortening.Key,
		OriginalURL: shortening.URL}); err != nil {
		return err
	}

	if err := refreshMap(&fs.m, config.FileStoragePath); err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) Get(identifier string) (*model.Shortening, error) {
	v, ok := fs.m.Load(identifier)
	if !ok {
		return nil, errors.New("not found")
	}

	shortening := v.(model.Shortening)

	return &shortening, nil
}

func refreshMap(m *sync.Map, filePath string) error {
	Reader, err := NewReader(filePath)
	if err != nil {
		return err
	}
	defer Reader.Close()

	files, err := Reader.ReadFile()
	if err != nil {
		return err
	}

	//erase SyncMap
	m.Range(func(key interface{}, value interface{}) bool {
		m.Delete(key)
		return true
	})

	for _, file := range files {
		m.Store(file.ShortURL,
			model.Shortening{
				Key: file.ShortURL,
				URL: file.OriginalURL})
	}

	return nil
}

func (fs *FileStorage) PutBatch(shorteningList []model.Shortening) error {
	for _, item := range shorteningList {
		err := fs.Put(item)
		if err != nil {
			return err
		}
	}
	return nil
}

// Initialize хранилища
func (fs *FileStorage) Initialize() error {
	return nil
}

// TODO
func (fs *FileStorage) GetByURL(url string) (*model.Shortening, error) {
	return nil, nil
}

// TODO
func (fs *FileStorage) GetByUserID(id string) ([]*model.Shortening, error) {
	return nil, nil
}
