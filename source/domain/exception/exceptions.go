package exception

import (
	"net/http"
)

type HttpException struct {
	StatusCode int
	Message    string
}

func NewNotFoundException(message string) HttpException {
	return HttpException{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}
