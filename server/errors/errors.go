package errors

import (
	"fmt"
	"net/http"
)

const (
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodeUnauthorized        = "UNAUTHORIZED"
	ErrCodeForbidden           = "FORBIDDEN"
	ErrCodeNotFound            = "NOT_FOUND"
	ErrCodeConflict            = "CONFLICT"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrGenericHttpError        = "HTTP_ERROR"

	ErrUserNotFound = "USER_NOT_FOUND"
	ErrTodoNotFound = "TODO_NOT_FOUND"
)

type ApiError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewBadRequest(message string) *ApiError {
	return &ApiError{
		Code:       ErrCodeBadRequest,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewUnauthorized(message string) *ApiError {
	return &ApiError{
		Code:       ErrCodeUnauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func NewForbidden(message string) *ApiError {
	return &ApiError{
		Code:       ErrCodeForbidden,
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		Err: err,
	}
}

func UserNotFound() *ApiError {
	return &ApiError{
		Code:       ErrUserNotFound,
		Message:    "User not found",
		StatusCode: http.StatusNotFound,
	}
}

func TodoNotFound() *ApiError {
	return &ApiError{
		Code:       ErrTodoNotFound,
		Message:    "Todo not found",
		StatusCode: http.StatusNotFound,
	}
}
