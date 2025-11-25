package exception

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(ctx echo.Context, err error) error {
	_, ok := err.(*ValidationError)
	if ok {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}	
	return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": err.Error(),
	})
}