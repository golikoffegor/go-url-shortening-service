package storage

// MemStorageInterface интерфейс
type MemStorageInterface interface {
	GetURLAddress(key string)
	UpdateURLAddress(key string, url string)
}

// MemStorage хранилище
type MemStorage struct {
	urlAddresses map[string]string
}

// NewMemStorage создает новый экземпляр MemStorage
func NewMemStorage() *MemStorage {
	return &MemStorage{
		urlAddresses: make(map[string]string),
	}
}

// GetURLAddress возвращает URL из хранилища по ключу key
func (m *MemStorage) GetURLAddress(key string) (string, bool) {
	v, ok := m.urlAddresses[key]
	return v, ok
}

// UpdateURLAddress записывает URL в хранилище с ключом key
func (m *MemStorage) UpdateURLAddress(key string, url string) {
	m.urlAddresses[key] = url
}
