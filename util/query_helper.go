package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetIntQuery retrieves an integer query parameter with a default value
func GetIntQuery(c echo.Context, key string, defaultValue int) int {
	str := c.QueryParam(key)
	if str == "" {
		return defaultValue
	}
	
	if val, err := strconv.Atoi(str); err == nil {
		return val
	}
	
	return defaultValue
}

// GetPaginationParams retrieves limit and offset from query parameters
func GetPaginationParams(c echo.Context) (limit int, offset int) {
	limit = GetIntQuery(c, "limit", 10)
	offset = GetIntQuery(c, "offset", 0)
	return
}
