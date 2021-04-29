package service

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"l6p.io/ui-api/pkg/cfg"
	"l6p.io/ui-api/pkg/util"
	"l6p.io/ui-api/pkg/v1/db"
	"l6p.io/ui-api/pkg/v1/vo"
	"math"
	"net/http"
)

func GetHttpKeyFactors(ctx echo.Context) error {
	conf := ctx.Get("config").(*cfg.Config)
	name := ctx.Param("p1")

	timelineData, err := db.FindHttpTimeline(conf, name)
	if err != nil {
		return ctx.NoContent(http.StatusOK)
	}

	summaryData, err := db.FindHttpSummary(conf, name)
	if err != nil {
		return ctx.NoContent(http.StatusOK)
	}

	if summaryData.Count == 0 {
		return ctx.NoContent(http.StatusOK)
	}

	var maxRpm = 0
	var sumRpm = 0
	var avgRpm = 0
	var count = 0

	for _, item := range timelineData {
		if maxRpm < item.Count {
			maxRpm = item.Count
		}
		sumRpm += item.Count
		count += 1
	}

	if count > 0 {
		avgRpm = int(math.Round(float64(sumRpm) / float64(count)))
	}

	var p95 = 0
	var p90 = 0
	var p75 = 0
	var p50 = 0
	for _, item := range summaryData.ByPercentile {
		switch item.Percentile {
		case 95:
			p95 = item.Duration
		case 90:
			p90 = item.Duration
		case 75:
			p75 = item.Duration
		case 50:
			p50 = item.Duration
		}
	}

	return ctx.JSON(http.StatusOK, vo.KeyFactors{
		MaxRpm:     maxRpm,
		AvgRpm:     avgRpm,
		HttpCount:  summaryData.Count,
		HttpFailed: summaryData.Failed,
		P95:        p95,
		P90:        p90,
		P75:        p75,
		P50:        p50,
	})
}

func GetHttpTimeline(ctx echo.Context) error {
	conf := ctx.Get("config").(*cfg.Config)
	tid := ctx.Param("p1")

	timelineData, err := db.FindHttpTimeline(conf, tid)
	if err != nil {
		return ctx.NoContent(http.StatusOK)
	}

	var startTime, warmUpDone, endTime int64
	summaryData, err := db.FindHttpSummary(conf, tid)
	if err == nil {
		startTime = summaryData.Start
		warmUpDone = summaryData.WarmUpDone
		endTime = summaryData.End
	} else if err == mgo.ErrNotFound {
		startTime = 0
		warmUpDone = 0
		endTime = 0
	} else {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, &struct {
		StartTime  int64       `json:"startTime"`
		WarmUpDone int64       `json:"warmUpDone"`
		EndTime    int64       `json:"endTime"`
		Timeline   interface{} `json:"timeline"`
	}{
		StartTime:  startTime,
		WarmUpDone: warmUpDone,
		EndTime:    endTime,
		Timeline:   timelineData,
	})
}

func GetHttpMethodAndStatus(ctx echo.Context) error {
	conf := ctx.Get("config").(*cfg.Config)
	tid := ctx.Param("p1")

	ret, err := db.FindHttpSummary(conf, tid)
	if err != nil {
		return ctx.NoContent(http.StatusOK)
	}

	if len(ret.ByMethod) == 0 || len(ret.ByStatus) == 0 {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, struct {
		ByMethod []*vo.HttpSummaryByMethodData `json:"byMethod"`
		ByStatus []*vo.HttpSummaryByStatusData `json:"byStatus"`
	}{
		ByMethod: ret.ByMethod,
		ByStatus: ret.ByStatus,
	})
}

func GetHttpUrls(ctx echo.Context) error {
	conf := ctx.Get("config").(*cfg.Config)
	tid := ctx.Param("p1")
	page := util.IntParam(ctx, "page")
	rowsPerPage := util.IntParam(ctx, "rowsPerPage")
	order := ctx.QueryParam("order")

	ret, err := db.FindHttpUrls(conf, tid, page, rowsPerPage, order)
	if err != nil {
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.JSON(http.StatusOK, ret)
}
