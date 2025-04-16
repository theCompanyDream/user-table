package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/theCompanyDream/user-table/apps/backend/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/ziflex/lecho"
	"golang.org/x/time/rate"

	"github.com/theCompanyDream/user-table/apps/backend/controller"
	"github.com/theCompanyDream/user-table/apps/backend/repository"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	server := echo.New()
	logger := lecho.New(
		os.Stdout,
		lecho.WithLevel(log.DEBUG),
		lecho.WithTimestamp(),
		lecho.WithCaller(),
	)

	server.HTTPErrorHandler = controller.HttpErrorHandler

	ulidController := controller.NewUlidController(db)
	uuid4Controller := controller.NewGormUuidController(db)
	nanoIdController := controller.NewGormNanoController(db)
	ksuidController := controller.NewGormKsuidController(db)
	cuidController := controller.NewGormCuidController(db)

	// Middleware
	server.Use(middleware.Recover())
	server.Logger = logger
	server.Use(middleware.RequestID())
	server.Use(middleware.Gzip())
	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))
	// Define main routes
	server.GET("/swagger/*", echoSwagger.WrapHandler)
	server.GET("/", controller.Home)
	server.GET("/ulids", ulidController.GetUsers)
	server.GET("/ulid/:id", ulidController.GetUser)
	server.POST("/ulid", ulidController.CreateUser)
	server.PUT("/ulid/:id", ulidController.UpdateUser)
	server.DELETE("/ulid/:id", ulidController.DeleteUser)
	//uuid
	server.GET("/uuid4", uuid4Controller.GetUsers)
	server.GET("/uuid4/:id", uuid4Controller.GetUser)
	server.POST("/uuid4", uuid4Controller.CreateUser)
	server.PUT("/uuid4/:id", uuid4Controller.UpdateUser)
	server.DELETE("/uuid4/:id", uuid4Controller.DeleteUser)
	//nanoId
	server.GET("/nano", nanoIdController.GetUsers)
	server.GET("/nano/:id", nanoIdController.GetUser)
	server.POST("/nano", nanoIdController.CreateUser)
	server.PUT("/nano/:id", nanoIdController.UpdateUser)
	server.DELETE("/nano/:id", nanoIdController.DeleteUser)
	//ksuid
	server.GET("/ksuid", ksuidController.GetUsers)
	server.GET("/ksuid/:id", ksuidController.GetUser)
	server.POST("/ksuid", ksuidController.CreateUser)
	server.PUT("/ksuid/:id", ksuidController.UpdateUser)
	server.DELETE("/ksuid/:id", ksuidController.DeleteUser)
	//cuid
	server.GET("/cuid", cuidController.GetUsers)
	server.GET("/cuid/:id", cuidController.GetUser)
	server.POST("/cuid", cuidController.CreateUser)
	server.PUT("/cuid/:id", cuidController.UpdateUser)
	server.DELETE("/cuid/:id", cuidController.DeleteUser)
	// Start the server
	server.Logger.Info("Server is running...")
	port := os.Getenv("BACKEND_PORT")
	if port != "" {
		serverStartCode := fmt.Sprintf(":%s", port)
		server.Logger.Fatal(server.Start(serverStartCode))
	} else {
		server.Logger.Fatal(server.Start(":3000"))
	}
}
