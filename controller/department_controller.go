package controller

import (
	"net/http"
	"packlib/model"
	"packlib/repository"
	"packlib/service"
	"packlib/util"

	"github.com/labstack/echo/v4"
)

type DepartmentController struct {
	departmentService service.DepartmentService
}

func NewDepartmentController(departmentService service.DepartmentService) DepartmentController {
	return DepartmentController{departmentService: departmentService}
}

func (controller *DepartmentController) Route(g *echo.Group) {
	g.GET("/departments", controller.GetDepartments)
}

func (controller *DepartmentController) GetDepartments(c echo.Context) error {
	limit, offset := util.GetPaginationParams(c)
	departments, err := controller.departmentService.Find(repository.DepartmentFindParams{
		BaseQueryParams: model.BaseQueryParams{
			Limit:  &limit,
			Offset: &offset,
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.WebResponse{
			Status:  "error",
			Message: "Failed to retrieve departments",
			Data:    nil,
		})
	}
	
	return c.JSON(http.StatusOK, model.WebResponse{
		Status:  "success",
		Message: "Departments retrieved successfully",
		Data:    departments,
	})
}

