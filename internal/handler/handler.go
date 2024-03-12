package handler

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/golikoffegor/go-server-metcrics-and-alerts/config"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/storage"
)

// Функция — обработчик HTTP-запроса RegistryHandlerURL
func RegistryHandlerURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	if len(body) > 0 && validateURL(string(body)) {
		// Генерируем ключ
		key := genShortURL()
		stringURL := string(body)
		// Сохраняем данные в хранилище
		storage.Storage.UpdateURLAddress(key, stringURL)
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		w.Write([]byte(config.BaseURL + "/" + key))
	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("No URL in request"))
	}
}

// Функция — обработчик HTTP-запроса GetURLbyIDHandler
func GetURLbyIDHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.String()[1:]
	getURL := storage.Storage.GetURLAddress(key)
	if getURL != "" {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", getURL)
		w.WriteHeader(307)
	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", "URL not found")
		w.WriteHeader(400)
	}
}

// Генерирование короткого URL
func genShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, 7)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}
	return string(key)
}

// Проверяем валидность URL
func validateURL(str string) bool {
	urlData, err := url.Parse(str)
	return err == nil && urlData.Host != "" && urlData.Scheme != ""
}
