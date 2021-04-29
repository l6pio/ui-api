package router

import (
	"github.com/labstack/echo/v4"
	"l6p.io/ui-api/pkg/v1/service"
)

func JobRouter(test *echo.Group) {
	test.GET("", func(ctx echo.Context) error { return service.GetJobs(ctx) })
	test.GET("/search/:p1", func(ctx echo.Context) error { return service.FindJobs(ctx) })
}
