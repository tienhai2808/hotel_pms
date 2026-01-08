package errors

import (
	"net/http"

	"github.com/InstayPMS/backend/pkg/constants"
)

var (
	ErrLoginFailed = NewAPIError(
		http.StatusBadRequest,
		constants.CodeLoginFailed,
		constants.SlugLoginFailed,
		"Incorrect username or password",
	)

	ErrInvalidToken = NewAPIError(
		http.StatusBadRequest,
		constants.CodeInvalidToken,
		constants.SlugInvalidToken,
		"Invalid or expired token",
	)
)

type APIError struct {
	Status  int
	Code    int
	Slug    string
	Message string
}

func NewAPIError(status, code int, slug, message string) *APIError {
	return &APIError{
		status,
		code,
		slug,
		message,
	}
}

func (e *APIError) Error() string {
	return e.Message
}
