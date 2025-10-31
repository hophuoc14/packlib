package middlewares

import "github.com/labstack/echo/v4"

func AnotherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Another middleware logic here
		return next(c)
	}
}