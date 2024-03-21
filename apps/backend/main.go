package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/theCompanyDream/user-angular/apps/backend/controller"
	"github.com/theCompanyDream/user-angular/apps/backend/repository"
)

func main() {
	repository.InitDB()
	server := echo.New()

	server.GET("/", controller.Home)
	server.POST("/user", controller.CreateUser)
	server.GET("/users", controller.GetUsers)
	server.GET("/user/:id", controller.GetUser)
	server.PUT("/user/:id", controller.UpdateUser)
	server.DELETE("/user/:id", controller.DeleteUser)

	server.Logger.Info("Server is running...")
	port := os.Getenv("BACKEND_PORT")
	if port != "" {
		serverStartCode := fmt.Sprintf(":%s", port)
		server.Logger.Fatal(server.Start(serverStartCode))
	} else {
		server.Logger.Fatal(server.Start(":3000"))
	}
}
