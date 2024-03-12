package config

import "flag"

var (
	ServerAdress string
	BaseURL      string
)

func ParseFlags() {
	flag.StringVar(&ServerAdress, "a", "localhost:8080", "Address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "Base URL for requests")
	flag.Parse()
}
