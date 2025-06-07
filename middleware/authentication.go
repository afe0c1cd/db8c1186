package middleware

import (
	"github.com/afe0c1cd/db8c1186/authn"
	"github.com/afe0c1cd/db8c1186/database"
	"github.com/afe0c1cd/db8c1186/server/errors"
	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(database database.Repository, authn authn.Repository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ah := c.Request().Header.Get("Authorization")
			if ah == "" {
				return errors.NewUnauthorized("authorization header is required")
			}
			if len(ah) < 7 || ah[:7] != "Bearer " {
				return errors.NewUnauthorized("invalid authorization header")
			}
			token := ah[7:]
			if token == "" {
				return errors.NewUnauthorized("token is required")
			}
			user, err := authn.AuthenticateByToken(token)
			if err != nil {
				return errors.NewUnauthorized("invalid token")
			}
			if user == nil {
				return errors.NewUnauthorized("user not found")
			}
			u, err := database.FindUserByID(c.Request().Context(), user.ID.String())
			if err != nil {
				return errors.NewInternalServerError(err)
			}
			if u == nil {
				return errors.UserNotFound()
			}

			c.Set("user", u)

			return next(c)
		}
	}
}
