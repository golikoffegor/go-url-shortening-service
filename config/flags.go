package config

import (
	"flag"
	"os"
)

var (
	ServerAdress string
	BaseURL      string
)

func ParseFlags() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		ServerAdress = envRunAddr
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		BaseURL = envBaseURL
	}

	flag.StringVar(&ServerAdress, "a", "localhost:8080", "Address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "Base URL for requests")
	flag.Parse()
}
