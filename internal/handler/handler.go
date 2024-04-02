package handler

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/golikoffegor/go-server-metcrics-and-alerts/config"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/model"
	"github.com/golikoffegor/go-server-metcrics-and-alerts/internal/storage"
)

var Storage *storage.MemStorage

// Функция — обработчик HTTP-запроса JsonHandlerURL
func JSONHandlerURL(w http.ResponseWriter, r *http.Request) {
	var inputJSON model.InputJSON
	body, err := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &inputJSON)
	if validateURL(inputJSON.URL) && err == nil {
		// Генерируем ключ
		key := genShortURL()
		// Сохраняем данные в хранилище
		Storage.UpdateURLAddress(key, inputJSON.URL)
		// Установка заголовков
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		resultModel := model.ResultShortenURL{URL: config.BaseURL + "/" + key}
		resultJSON, _ := json.Marshal(resultModel)
		_, err := w.Write(resultJSON)
		if err != nil {
			return
		}

	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		_, err := w.Write([]byte("No URL in request"))
		if err != nil {
			return
		}
	}
}

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
		Storage.UpdateURLAddress(key, stringURL)
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		_, err := w.Write([]byte(config.BaseURL + "/" + key))
		if err != nil {
			return
		}

	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		_, err := w.Write([]byte("No URL in request"))
		if err != nil {
			return
		}
	}
}

// Функция — обработчик HTTP-запроса GetURLbyIDHandler
func GetURLbyIDHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.String()[1:]
	getURL, ok := Storage.GetURLAddress(key)
	okURL := validateURL(getURL)
	if ok && okURL {
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
