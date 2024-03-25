package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
* @Success 200 {200} string "ok"
**/
func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
