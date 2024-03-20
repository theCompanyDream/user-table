package main

import (
    "database/sql"
    "fmt"
    "log"
	"github.com/labstack/echo/v4"
    _ "github.com/lib/pq"

    "github.com/theCompanyDream/user-angular/apps/backend/controller"
)

func main() {
    initDB()
    defer db.Close()
    server := echo.New()

    e.POST("/user", createUser)
    e.GET("/user/:id", getUser)
    e.PUT("/user/:id", updateUser)
    e.DELETE("/user/:id", deleteUser)

    e.Logger.Info("Server is running...")
}
