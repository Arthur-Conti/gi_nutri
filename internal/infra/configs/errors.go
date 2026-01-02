package configs

import (
	"fmt"
	"net/http"
)

// AppError representa um erro da aplicação com código HTTP
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError cria um novo erro da aplicação
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Erros comuns
var (
	ErrNotFound          = NewAppError(http.StatusNotFound, "resource not found", nil)
	ErrInvalidID         = NewAppError(http.StatusBadRequest, "invalid id format", nil)
	ErrInvalidInput      = NewAppError(http.StatusBadRequest, "invalid input data", nil)
	ErrInternalServer     = NewAppError(http.StatusInternalServerError, "internal server error", nil)
	ErrPatientNotFound   = NewAppError(http.StatusNotFound, "patient not found", nil)
	ErrResultsNotFound   = NewAppError(http.StatusNotFound, "results not found", nil)
	ErrValidationFailed  = NewAppError(http.StatusBadRequest, "validation failed", nil)
)

// WrapError envolve um erro existente com um AppError
func WrapError(err error, appErr *AppError) *AppError {
	return NewAppError(appErr.Code, appErr.Message, err)
}

