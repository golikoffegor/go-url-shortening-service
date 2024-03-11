package server

import (
	"net/http"

	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/handler"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/storage"
)

// Инициализации зависимостей сервера перед запуском
func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routeMethod)
	return http.ListenAndServe(`:8080`, mux)
}

// Метод распределения запросов по типу
func routeMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.RegistryHandlerURL(w, r)
	case http.MethodGet:
		handler.GetURLbyIDHandler(w, r)
	}
}

// Инициализация и запуск сервера
func InitAndRunServer() {
	storage.Storage = storage.NewMemStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("/", routeMethod)
	if err := run(); err != nil {
		panic(err)
	}
}
