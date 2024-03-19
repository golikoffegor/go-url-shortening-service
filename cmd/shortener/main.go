package main

import (
	"github.com/golikoffegor/go-server-metcrics-and-alerts/config"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/server"
)

// Функция main вызывается автоматически при запуске приложения
func main() {
	config.ParseFlags()
	server.InitAndRunServer()
}
