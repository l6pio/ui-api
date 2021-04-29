package service

import (
	"github.com/labstack/echo/v4"
	"l6p.io/ui-api/pkg/cfg"
	"l6p.io/ui-api/pkg/v1/db"
	"net/http"
)

func GetJobs(ctx echo.Context) error {
	conf := ctx.Get("config").(*cfg.Config)

	ret, err := db.GetJobs(conf)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, ret)
}

func FindJobs(ctx echo.Context) error {
	conf := ctx.Get("config").(*cfg.Config)
	keyword := ctx.Param("p1")

	ret, err := db.FindJobsByKeyword(conf, keyword)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, ret)
}
