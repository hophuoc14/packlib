package controller

import (
	"net/http"
	"packlib/model"
	"packlib/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}

func (controller *UserController) Route(app *echo.Echo) {
	app.POST("/register", controller.Register)
	app.POST("/login", controller.Login)
}

func (controller *UserController) Register(c echo.Context) error {
	var request model.CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "error",
			Message: "Bad Request",
			Data:    err,
		})
	}
	response, err := controller.userService.Register(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "error",
			Message: "Bad Request",
			Data:    err,
		})
	}
	return c.JSON(http.StatusOK, model.WebResponse{
		Status:  "success",
		Message: "Register Success",
		Data:    response,
	})
}

func (controller *UserController) Login(c echo.Context) error {
	var request model.CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "error",
			Message: "Bad Request",
			Data:    err,
		})
	}
	response, err := controller.userService.Login(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "error",
			Message: "Bad Request",
			Data:    err,
		})
	}
	return c.JSON(http.StatusOK, model.WebResponse{
		Status:  "success",
		Message: "Login Success",
		Data:    response,
	})
}