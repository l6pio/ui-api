package middlewares

import (
	"github.com/labstack/echo/v4"
	"l6p.io/ui-api/pkg/cfg"
)

func Config(conf *cfg.Config) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		setConfig(conf),
	}
}

func setConfig(conf *cfg.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("config", conf)
			return next(ctx)
		}
	}
}
