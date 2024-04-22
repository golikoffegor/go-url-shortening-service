package model

type InputURLData struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type OutputURLData struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
