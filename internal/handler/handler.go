package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/golikoffegor/go-url-shortening-service/config"
	"github.com/golikoffegor/go-url-shortening-service/internal/interfaces"
	"github.com/golikoffegor/go-url-shortening-service/internal/model"
)

var Storage interfaces.Storager

// Функция — обработчик HTTP-запроса JsonHandlerURL
func JSONHandlerURL(w http.ResponseWriter, r *http.Request) {
	var inputJSON model.InputJSON
	body, err := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &inputJSON)
	if validateURL(inputJSON.URL) && err == nil {
		// Генерируем ключ
		shortening := model.Shortening{Key: genShortURL(), URL: inputJSON.URL}
		// Сохраняем данные в хранилище
		err := Storage.Put(shortening)
		if err != nil {
			return
		}
		// Установка заголовков
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		resultModel := model.ResultShortenURL{URL: config.BaseURL + "/" + shortening.Key}
		resultJSON, _ := json.Marshal(resultModel)
		_, err = w.Write(resultJSON)
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
func BatchRegistryHandlerURL(w http.ResponseWriter, r *http.Request) {
	var inputJSON []model.InputURLData
	body, err := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &inputJSON)
	if validateURLs(inputJSON) && err == nil {
		var urlBatch []model.Shortening
		var outputURLBatch []model.OutputURLData
		for _, item := range inputJSON {
			// Генерируем ключ
			key := genShortURL()
			urlBatch = append(urlBatch, model.Shortening{Key: key, URL: item.OriginalURL})
			outputURLBatch = append(outputURLBatch, model.OutputURLData{CorrelationID: item.CorrelationID, ShortURL: config.BaseURL + "/" + key})
		}
		// Сохраняем данные в хранилище
		err := Storage.PutBatch(urlBatch)
		if err != nil {
			return
		}
		// Установка заголовков
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		resultJSON, _ := json.Marshal(outputURLBatch)
		_, err = w.Write(resultJSON)
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
		shortening := model.Shortening{Key: genShortURL(), URL: string(body)}
		// Сохраняем данные в хранилище
		err := Storage.Put(shortening)
		if err != nil {
			return
		}
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		_, err = w.Write([]byte(config.BaseURL + "/" + shortening.Key))
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
	shortening, err := Storage.Get(key)
	okURL := validateURL(shortening.URL)
	if err == nil && okURL {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", shortening.URL)
		w.WriteHeader(307)
	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", "URL not found")
		w.WriteHeader(400)
	}
}

// Функция — обработчик HTTP-запроса GetPingDB
func GetPingDB(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("pgx", config.PostgreSQLDSN)
	err := db.Ping()
	if err == nil {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
	} else {
		// Установка заголовков
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
	}
	defer db.Close()
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

// Проверяем валидность URL
func validateURLs(InputURLData []model.InputURLData) bool {
	for _, item := range InputURLData {
		if !validateURL(item.OriginalURL) {
			return false
		}
	}
	return true
}
