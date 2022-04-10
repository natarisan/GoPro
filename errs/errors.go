package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

//404
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusNotFound,
	}
}

//500
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusInternalServerError,
	}
}

func(e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusUnprocessableEntity,
	}
}