package main

import (
	"packlib/config"
	"packlib/controller"
	"packlib/exception"
	"packlib/repository"
	"packlib/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configuration := config.New()
	database, err := config.NewPostgresDatabase(configuration)
	
	if err != nil {
		panic(err)
	}
	
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	
	departmentRepository := repository.NewDepartmentRepository(database)
	departmentService := service.NewDepartmentService(departmentRepository)
	departmentController := controller.NewDepartmentController(departmentService)
	
	e := echo.New()
	api := e.Group("/api")

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		err = exception.ErrorHandler(context, err)
	}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH , echo.DELETE},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("10MB"))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// routing
	userController.Route(api)
	departmentController.Route(api)	
	// start server
	e.Logger.Fatal(e.Start(":8080"))

}