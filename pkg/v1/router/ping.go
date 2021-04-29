package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PingRouter(ping *echo.Group) {
	ping.GET("", func(ctx echo.Context) error { return ctx.String(http.StatusOK, "pong") })
}
