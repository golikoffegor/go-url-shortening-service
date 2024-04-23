package model

type ConflictError struct {
}

func (ce ConflictError) Error() string {
	return ""
}
