package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golikoffegor/go-url-shortening-service/config"
	"github.com/golikoffegor/go-url-shortening-service/internal/model"
	"github.com/golikoffegor/go-url-shortening-service/internal/storage/memstorage"
)

func TestJSONHandlerURL(t *testing.T) {
	type request struct {
		method      string
		body        string
		contentType string
	}
	type want struct {
		code        int
		contentType string
		contentBody string
	}
	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "Test StatusCreated",
			request: request{
				method:      http.MethodPost,
				body:        `{"URL":"http://practicum.yandex.ru/"}`,
				contentType: "application/json",
			},
			want: want{
				code:        http.StatusCreated,
				contentType: "application/json",
				contentBody: `{"result":"`,
			},
		},
		{
			name: "Test StatusBadRequest",
			request: request{
				method:      http.MethodPost,
				body:        "",
				contentType: "application/json",
			},
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain",
				contentBody: "No URL in request",
			},
		},
	}
	config.BaseURL = "http://localhost:8080"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Создаем фейковое хранилище
			Storage = memstorage.NewStorage()
			// Создаем HTTP запрос для записи URL
			req := httptest.NewRequest(test.request.method, "/api/shorten", strings.NewReader(test.request.body))
			// Задаем заголовки
			req.Header.Set("Content-Type", test.request.contentType)
			// Создаем ResponseRecorder для записи ответа сервера
			w := httptest.NewRecorder()
			// Создаем обработчик и вызываем его метод ServeHTTP
			JSONHandlerURL(w, req)
			// Результат записываем в переменную
			result := w.Result()
			defer result.Body.Close()
			// Проверяем код ответа
			assert.Equal(t, test.want.code, result.StatusCode)
			// Проверяем заголовки ответа
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))
			// Проверяем тело ответа запроса
			assert.Contains(t, w.Body.String(), test.want.contentBody)
		})
	}
}

func TestRegistryHandlerURL(t *testing.T) {
	type request struct {
		method      string
		body        string
		contentType string
	}
	type want struct {
		code        int
		contentType string
		contentBody string
	}
	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "Test StatusCreated",
			request: request{
				method:      http.MethodPost,
				body:        "https://practicum.yandex.ru/",
				contentType: "text/plain",
			},
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
				contentBody: "http://localhost:8080/",
			},
		},
		{
			name: "Test StatusBadRequest",
			request: request{
				method:      http.MethodPost,
				body:        "",
				contentType: "text/plain",
			},
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain",
				contentBody: "No URL in request",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Создаем фейковое хранилище
			Storage = memstorage.NewStorage()
			// Создаем HTTP запрос для записи URL
			req := httptest.NewRequest(test.request.method, "/", strings.NewReader(test.request.body))
			// Задаем заголовки
			req.Header.Set("Content-Type", test.request.contentType)
			// Создаем ResponseRecorder для записи ответа сервера
			w := httptest.NewRecorder()
			// Создаем обработчик и вызываем его метод ServeHTTP
			RegistryHandlerURL(w, req)
			// Результат записываем в переменную
			result := w.Result()
			defer result.Body.Close()
			// Проверяем код ответа
			assert.Equal(t, test.want.code, result.StatusCode)
			// Проверяем заголовки ответа
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))
			// Проверяем тело ответа запроса
			assert.Contains(t, "http://localhost:8080/"+w.Body.String(), test.want.contentBody)

		})
	}
}

func TestGetURLbyIDHandler(t *testing.T) {
	type request struct {
		method      string
		urlAddr     string
		contentType string
	}
	type want struct {
		code        int
		contentType string
		location    string
	}
	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "Test StatusTemporaryRedirect",
			request: request{
				method:      http.MethodGet,
				urlAddr:     "http://localhost:8080/",
				contentType: "text/plain",
			},
			want: want{
				code:        http.StatusTemporaryRedirect,
				contentType: "text/plain",
				location:    "https://practicum.yandex.ru/",
			},
		},
		{
			name: "Test StatusBadRequest",
			request: request{
				method:      http.MethodGet,
				urlAddr:     "http://localhost:8080/",
				contentType: "text/plain",
			},
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain",
				location:    "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Создаем фейковое хранилище
			Storage = memstorage.NewStorage()
			shortening := model.Shortening{Key: genShortURL(), URL: test.want.location}
			_ = Storage.Put(shortening)
			// Создаем HTTP запрос для записи URL
			req := httptest.NewRequest(test.request.method, "/"+shortening.Key, strings.NewReader(""))
			// Задаем заголовки
			req.Header.Set("Content-Type", test.request.contentType)
			// Создаем ResponseRecorder для записи ответа сервера
			w := httptest.NewRecorder()
			// Создаем обработчик и вызываем его метод ServeHTTP
			GetURLbyIDHandler(w, req)
			// Результат записываем в переменную
			result := w.Result()
			defer result.Body.Close()
			// Проверяем код ответа
			assert.Equal(t, test.want.code, result.StatusCode)
			// Проверяем заголовки ответа
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))
			// Проверяем тело ответа запроса
			if test.want.location != "" {
				assert.Contains(t, result.Header.Get("location"), test.want.location)
			}
		})
	}
}
