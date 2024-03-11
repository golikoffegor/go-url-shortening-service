// Пакеты исполняемых приложений должны называться main
package main

import (
	"fmt"
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
	http.HandleFunc("/", registerUrl)
	http.HandleFunc("/{id}", getUrlById)
	return http.ListenAndServe(`:8080`, nil)
}

// Функция — обработчик HTTP-запроса getUrlById
func getUrlById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Разрешаем только GET-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
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

// Функция — обработчик HTTP-запроса registerUrl
func registerUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Разрешаем только POST-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 || !isURL(string(body)) {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", "URL not found")
		w.WriteHeader(400)
	}

	stringUrl := string(body)
	key := saveUrl(stringUrl)
	fmt.Println(urlAddresses[key])

	// Установка заголовков
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", "30")
	w.WriteHeader(201)
	_, _ = w.Write([]byte("https://localhost:8080/" + key))
}

// Сохранение URL
func saveUrl(url string) string {
	key := genShortUrl()
	urlAddresses[key] = url
	return key
}

// Генерирование короткого URL
func genShortUrl() string {
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
