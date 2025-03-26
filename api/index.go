package handler

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/theCompanyDream/user-angular/apps/backend/controller"
	"github.com/theCompanyDream/user-angular/apps/backend/repository"
)

// Handler is the AWS Lambda handler function.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize DB only once
	initDBOnce.Do(repository.InitDB)

	// Create a new Echo server
	server := echo.New()

	// CORS middleware for Vercel
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

	// Adjust routes to include /api prefix
	server.GET("/api", controller.GetUsers)
	server.GET("/api/users", controller.GetUsers)
	server.GET("/api/user/:id", controller.GetUser)
	server.POST("/api/user", controller.CreateUser)
	server.PUT("/api/user/:id", controller.UpdateUser)
	server.DELETE("/api/user/:id", controller.DeleteUser)

	server.ServeHTTP(w, r)
}

// Add this to ensure DB is initialized only once
var initDBOnce sync.Once
