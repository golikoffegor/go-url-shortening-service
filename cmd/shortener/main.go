package main

import (
	"github.com/golikoffegor/go-url-shortening-service/config"
	"github.com/golikoffegor/go-url-shortening-service/internal/server"
)

// Функция main вызывается автоматически при запуске приложения
func main() {
	config.ParseFlags()
	server.InitAndRunServer()
}
