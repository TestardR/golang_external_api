package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	APIStatus  int    `json:"status"`
	APIMessage string `json:"message"`
	APIError   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.APIStatus
}

func (e *apiError) Message() string {
	return e.APIMessage
}

func (e *apiError) Error() string {
	return e.APIError
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		APIStatus:  statusCode,
		APIMessage: message,
	}
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		APIStatus:  http.StatusNotFound,
		APIMessage: message,
	}
}

func NewInternalServer(message string) ApiError {
	return &apiError{
		APIStatus:  http.StatusInternalServerError,
		APIMessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		APIStatus:  http.StatusBadRequest,
		APIMessage: message,
	}
}

func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiError

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}

	return &result, nil
}
