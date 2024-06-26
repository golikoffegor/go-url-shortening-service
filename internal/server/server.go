package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/golikoffegor/go-url-shortening-service/config"
	"github.com/golikoffegor/go-url-shortening-service/internal/handler"
	"github.com/golikoffegor/go-url-shortening-service/internal/middleware"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage"
)

// Инициализации зависимостей сервера перед запуском
func run() error {
	r := chi.NewRouter()
	r.Post("/api/shorten", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.JSONHandlerURL))))
	r.Post("/api/shorten/batch", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.BatchRegistryHandlerURL))))
	r.Post("/", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.RegistryHandlerURL))))
	r.Get("/{id}", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.GetURLbyIDHandler))))
	r.Get("/ping", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.GetPingDB))))
	r.Get("/api/user/urls", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.GetByUserID))))
	r.Delete("/api/user/urls", middleware.WithAuth(middleware.GZIP(middleware.Log(handler.DeleteByUserID))))
	fmt.Println("Running server on", config.ServerAdress)
	return http.ListenAndServe(config.ServerAdress, r)
}

// Инициализация и запуск сервера
func InitAndRunServer() {
	handler.Storage = storage.GetStorage()
	err := handler.Storage.Initialize()
	if err != nil {
		panic(err)
	}
	if err := run(); err != nil {
		panic(err)
	}
}
