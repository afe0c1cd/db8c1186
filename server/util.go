package server

import (
	"github.com/afe0c1cd/db8c1186/model"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) *model.User {
	user, ok := c.Get("user").(*model.User)
	if !ok {
		return nil
	}
	return user
}
