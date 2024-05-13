package model

type InputJSON struct {
	URL string `json:"url"`
}

type ResultShortenURL struct {
	URL string `json:"result"`
}

type ResultShortenData struct {
	ShortURL string `json:"short_url"`
	URL      string `json:"original_url"`
}
