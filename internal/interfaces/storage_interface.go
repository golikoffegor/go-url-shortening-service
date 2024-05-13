package interfaces

import "github.com/golikoffegor/go-url-shortening-service/internal/model"

type Storager interface {
	Put(shortening model.Shortening) error
	GetByURL(url string) (*model.Shortening, error)
	GetByUserID(id string) ([]*model.Shortening, error)
	Get(identifier string) (*model.Shortening, error)
	PutBatch(shorteningList []model.Shortening) error
	Initialize() error
}
