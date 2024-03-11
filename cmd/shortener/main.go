// Пакеты исполняемых приложений должны называться main
package main

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"
)

var urlAddresses map[string]string

// Функция main вызывается автоматически при запуске приложения
func main() {
	urlAddresses = make(map[string]string)
	if err := run(); err != nil {
		panic(err)
	}
}

// Функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	http.HandleFunc("/", registerURL)
	http.HandleFunc("/{id}", getURLByID)
	return http.ListenAndServe(`:8080`, nil)
}

// Функция — обработчик HTTP-запроса getURLByID
func getURLByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Разрешаем только GET-запросы
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := r.URL.String()[1:]

	if value, ok := urlAddresses[key]; ok {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", value)
		w.WriteHeader(307)
	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", "URL not found")
		w.WriteHeader(400)
	}
}

// Функция — обработчик HTTP-запроса registerURL
func registerURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Разрешаем только POST-запросы
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 || !isURL(string(body)) {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", "URL not found")
		w.WriteHeader(400)
	}

	stringURL := string(body)
	key := saveURL(stringURL)

	// Установка заголовков
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", "30")
	w.WriteHeader(201)
	_, _ = w.Write([]byte("http://localhost:8080/" + key))
}

// Сохранение URL
func saveURL(url string) string {
	key := genShortURL()
	urlAddresses[key] = url
	return key
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
func isURL(str string) bool {
	urlData, err := url.Parse(str)
	return err == nil && urlData.Scheme != "" && urlData.Host != ""
}
