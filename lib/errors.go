package lib

import "net/http"

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
}

// Error implements error.
func (err *ValidationError) Error() string {
	return "VALIDATION_ERROR"
}

const (
	ErrorCodeInvalid        = 10
	ErrorCodeNotFound       = 11
	ErrorCodeDuplicateEntry = 12
	ErrorCodeNotAllowed     = 13
)

var DatabaseErrorCodeMappings = map[int]int{
	ErrorCodeInvalid:        http.StatusBadRequest,
	ErrorCodeNotFound:       http.StatusNotFound,
	ErrorCodeDuplicateEntry: http.StatusBadRequest,
	ErrorCodeNotAllowed:     http.StatusBadRequest,
}

type DatabaseError struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

// Error implements error.
func (err DatabaseError) Error() string {
	return err.Message
}

type HttpError struct {
	Message    interface{} `json:"message"`
	StatusCode int         `json:"statusCode"`
}

// Error implements error.
func (err HttpError) Error() string {
	return err.Message.(string)
}
