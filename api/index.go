package handler

import (
	"github.com/aws/aws-lambda-go/events"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/theCompanyDream/user-angular/apps/backend/controller"
	"github.com/theCompanyDream/user-angular/apps/backend/repository"
)

var echoLambda *echoadapter.EchoLambda

// initServer initializes the Echo server and sets up routes.
func initServer() *echo.Echo {
	// Initialize the database
	if err := repository.InitDB(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

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

// Handler is the AWS Lambda handler function.
func Index(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Initialize the Echo server
	server := initServer()

	// Wrap the Echo server with the adapter
	echoLambda = echoadapter.New(server)

	// Proxy the incoming request to the Echo server
	return echoLambda.Proxy(req)
}
