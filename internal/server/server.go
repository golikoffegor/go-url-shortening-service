package server

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/handler"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/storage"
)

// Инициализации зависимостей сервера перед запуском
func run() error {
	r := chi.NewRouter()
	r.Post("/", handler.RegistryHandlerURL)
	r.Get("/{id}", handler.GetURLbyIDHandler)
	return http.ListenAndServe(":8080", r)
}

// Инициализация и запуск сервера
func InitAndRunServer() {
	storage.Storage = storage.NewMemStorage()
	if err := run(); err != nil {
		panic(err)
	}
}
