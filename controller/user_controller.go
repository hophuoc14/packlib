package controller

import (
	"fmt"
	"net/http"
	"packlib/model"
	"packlib/service"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}

func (controller *UserController) Route(g *echo.Group) {
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
}

func (controller *UserController) Register(c echo.Context) error {
	var request model.CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "error",
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	response, err := controller.userService.Register(request)
	if err != nil {
		// Determine appropriate status code based on error type
		statusCode := http.StatusBadRequest
		message := "Registration failed"
		
		// Check for specific error types
		errMsg := err.Error()
		if strings.Contains(errMsg, "already exists") || strings.Contains(errMsg, "duplicate") {
			statusCode = http.StatusConflict
			message = "User already exists"
		} else if strings.Contains(errMsg, "validation") || strings.Contains(errMsg, "invalid") {
			statusCode = http.StatusBadRequest
			message = "Invalid input data"
		} else {
			statusCode = http.StatusInternalServerError
			message = "Internal server error"
		}

		return c.JSON(statusCode, model.WebResponse{
			Status:  "error",
			Message: message,
			Data:    nil, // Don't expose error details to client
		})
	}

	return c.JSON(http.StatusCreated, model.WebResponse{
		Status:  "success",
		Message: "Registration successful",
		Data:    response,
	})
}

func (controller *UserController) Login(c echo.Context) error {
	var request model.LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "error",
			Message: "Invalid request format",
			Data:    nil,
		})
	}
	fmt.Println(request, "come???")

	response, err := controller.userService.Login(request)
	if err != nil {
		// Determine appropriate status code based on error type
		statusCode := http.StatusUnauthorized
		message := "Invalid credentials"
		
		errMsg := err.Error()
		if strings.Contains(errMsg, "not found") {
			statusCode = http.StatusUnauthorized
			message = "Invalid username or password"
		} else if strings.Contains(errMsg, "password") {
			statusCode = http.StatusUnauthorized
			message = "Invalid username or password"
		} else if strings.Contains(errMsg, "validation") {
			statusCode = http.StatusBadRequest
			message = "Invalid input data"
		} else {
			statusCode = http.StatusInternalServerError
			message = "Internal server error"
		}

		return c.JSON(statusCode, model.WebResponse{
			Status:  "error",
			Message: message,
			Data:    nil, // Don't expose error details to client
		})
	}

	return c.JSON(http.StatusOK, model.WebResponse{
		Status:  "success",
		Message: "Login successful",
		Data:    response,
	})
}