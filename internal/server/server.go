package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/golikoffegor/go-server-metcrics-and-alerts/config"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/handler"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/storage"
)

// Инициализации зависимостей сервера перед запуском
func run() error {
	r := chi.NewRouter()
	r.Post("/", handler.RegistryHandlerURL)
	r.Get("/{id}", handler.GetURLbyIDHandler)
	fmt.Println("Running server on", config.ServerAdress)
	return http.ListenAndServe(config.ServerAdress, r)
}

// Инициализация и запуск сервера
func InitAndRunServer() {
	handler.Storage = storage.NewMemStorage()
	if err := run(); err != nil {
		panic(err)
	}

}
