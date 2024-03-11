package storage

var Storage *MemStorage

// urlType определяет тип URL
type urlType string

// MetricsStorage определяет интерфейс для взаимодействия с хранилищем метрик
type MetricsStorage interface {
	GetURLAddress(key string) string
	UpdateURLAddress(key string, url string) map[urlType]map[string]interface{}
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
func (m *MemStorage) GetURLAddress(key string) string {
	return m.urlAddresses[key]
}

// UpdateURLAddress записывает URL в хранилище с ключом key
func (m *MemStorage) UpdateURLAddress(key string, url string) {
	m.urlAddresses[key] = url
}
