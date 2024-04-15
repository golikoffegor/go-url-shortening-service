package interfaces

import "github.com/golikoffegor/go-url-shortening-service/internal/model"

type Storager interface {
	Put(shortening model.Shortening) error
	Get(identifier string) (*model.Shortening, error)
	Initialize() error
}
