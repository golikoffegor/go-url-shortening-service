package config

import (
	"flag"
	"os"
)

var (
	ServerAdress    string
	BaseURL         string
	FileStoragePath string
	PostgreSQLDSN   string
)

func ParseFlags() {
	flag.StringVar(&ServerAdress, "a", "localhost:8080", "Address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "Base URL for requests")
	flag.StringVar(&FileStoragePath, "f", "/tmp/short-url-db.json",
		"full name of the file where data in JSON format is saved")
	flag.StringVar(&PostgreSQLDSN, "d", "postgres://postgres:postgres@127.0.0.1:5432/", "Base DSN for PostgreSQL")
	flag.Parse()
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		ServerAdress = envRunAddr
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		BaseURL = envBaseURL
	}
	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		FileStoragePath = envFileStoragePath
	}
	if envPostgreSQLDSN := os.Getenv("DATABASE_DSN"); envPostgreSQLDSN != "" {
		PostgreSQLDSN = envPostgreSQLDSN
	}
}
