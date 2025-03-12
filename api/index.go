package handler

import (
	"net/http"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/theCompanyDream/user-angular/apps/backend/controller"
	"github.com/theCompanyDream/user-angular/apps/backend/repository"
)

var echoLambda *echoadapter.EchoLambda

// Handler is the AWS Lambda handler function.
func HandlerHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the Echo server
	repository.InitDB()

	// Create a new Echo server
	server := echo.New()

	// Middleware
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

	// Routes
	server.GET("/", controller.Home)
	server.GET("/users", controller.GetUsers)
	server.GET("/user/:id", controller.GetUser)
	server.POST("/user", controller.CreateUser)
	server.PUT("/user/:id", controller.UpdateUser)
	server.DELETE("/user/:id", controller.DeleteUser)

	return server
}
