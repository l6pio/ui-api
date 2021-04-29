package router

import (
	"github.com/labstack/echo/v4"
	"l6p.io/ui-api/pkg/v1/service"
)

func ReportRouter(report *echo.Group) {
	report.GET("/:p1/http/key-factors", func(ctx echo.Context) error { return service.GetHttpKeyFactors(ctx) })
	report.GET("/:p1/http/timeline", func(ctx echo.Context) error { return service.GetHttpTimeline(ctx) })
	report.GET("/:p1/http/method-and-status", func(ctx echo.Context) error { return service.GetHttpMethodAndStatus(ctx) })
	report.GET("/:p1/http/url", func(ctx echo.Context) error { return service.GetHttpUrls(ctx) })
}
