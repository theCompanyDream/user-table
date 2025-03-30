package handler

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/theCompanyDream/user-angular/apps/backend/controller"
	"github.com/theCompanyDream/user-angular/apps/backend/repository"
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
	e.GET("/api", controller.GetUsers)
	e.GET("/api/users", controller.GetUsers)
	e.GET("/api/user/:id", controller.GetUser)
	e.POST("/api/user", controller.CreateUser)
	e.PUT("/api/user/:id", controller.UpdateUser)
	e.DELETE("/api/user/:id", controller.DeleteUser)
}

// Handler is the AWS Lambda handler function.
func Handler(w http.ResponseWriter, r *http.Request) {
	initDBOnce.Do(repository.InitDB)
	echoOnce.Do(initEcho)

	e.ServeHTTP(w, r)
}
