package handler

import (
	"net/http"
	"sync"
	"fmt"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/theCompanyDream/user-table/apps/backend/controller"
	"github.com/theCompanyDream/user-table/apps/backend/repository"
)

// Echo instance
var (
	e          *echo.Echo
	echoOnce   sync.Once
	initDBOnce sync.Once
)

func initEcho() {
	e = echo.New()

	// Middleware setup
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Adjust routes to include /api prefix
	e.GET("/api", controller.Home)
	e.GET("/api/users", controller.GetUsers)
	e.GET("/api/user/:id", controller.GetUser)
	e.POST("/api/user", controller.CreateUser)
	e.PUT("/api/user/:id", controller.UpdateUser)
	e.DELETE("/api/user/:id", controller.DeleteUser)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize the database, capturing errors with full stack trace.
	initDBOnce.Do(func() {
		if err := repository.ServerlessInitDB(); err != nil {
			http.Error(w, fmt.Sprintf("Database initialization error: %v\n%s", err, debug.Stack()), http.StatusInternalServerError)
			return
		}
	})
	echoOnce.Do(initEcho)

	// Pass the request to Echo's HTTP handler.
	e.ServeHTTP(w, r)
}