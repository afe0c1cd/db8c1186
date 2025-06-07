package middleware

import (
	"errors"
	"net/http"

	"github.com/afe0c1cd/db8c1186/generated"
	herrors "github.com/afe0c1cd/db8c1186/server/errors"
	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, ctx echo.Context) {
	code := herrors.ErrCodeInternalServerError
	message := "Internal Server Error"
	status := http.StatusInternalServerError

	var apiErr *herrors.ApiError
	var httpErr *echo.HTTPError
	var internalErr *herrors.InternalServerError

	switch {
	case errors.As(err, &internalErr):
		// noop
	case errors.As(err, &apiErr):
		code = apiErr.Code
		message = apiErr.Message
		status = apiErr.StatusCode
	case errors.As(err, &httpErr):
		status = httpErr.Code
		switch status {
		case http.StatusBadRequest:
			code = herrors.ErrCodeBadRequest
		case http.StatusUnauthorized:
			code = herrors.ErrCodeUnauthorized
		case http.StatusForbidden:
			code = herrors.ErrCodeForbidden
		case http.StatusNotFound:
			code = herrors.ErrCodeNotFound
		case http.StatusConflict:
			code = herrors.ErrCodeConflict
		default:
			code = herrors.ErrGenericHttpError
		}
		if httpErr.Message != "" {
			message, _ = httpErr.Message.(string)
		} else {
			message = http.StatusText(status)
		}
	}

	ctx.JSON(status, generated.ErrorResponse{
		Code:    code,
		Message: message,
	})
}
