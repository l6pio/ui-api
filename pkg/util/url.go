package util

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func IntParam(ctx echo.Context, name string) int {
	param := ctx.QueryParam(name)

	ret := 0
	if param != "" {
		ret, _ = strconv.Atoi(param)
	}
	return ret
}
