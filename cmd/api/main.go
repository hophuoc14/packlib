package main

import (
	"fmt"
	"os"
	"packlib/cmd/api/handlers"
	"packlib/common"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	logger echo.Logger
	server *echo.Echo 
	handler handlers.Handler
}

func main() {
	e := echo.New()
	// load .env file
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}
	// connect to db
	db, err := common.NewPgDB()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	_ = db
	// initialize application
	handler := handlers.Handler{
		DB: db,
	}
	app := Application{
		logger: e.Logger,
		server: e,
		handler: handler,
	}
	e.Use(middleware.Logger())

app.routes()
	fmt.Println(app)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}