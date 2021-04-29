package vo

import "reflect"

type KeyFactors struct {
	MaxRpm     int `json:"maxRpm"`
	AvgRpm     int `json:"avgRpm"`
	HttpCount  int `json:"httpCount"`
	HttpFailed int `json:"httpFailed"`
	P95        int `json:"p95"`
	P90        int `json:"p90"`
	P75        int `json:"p75"`
	P50        int `json:"p50"`
}

type HttpTimeline struct {
	From   int64 `json:"from" bson:"from"`
	To     int64 `json:"to" bson:"to"`
	Count  int   `json:"count" bson:"count"`
	Failed int   `json:"failed" bson:"failed"`
}

type HttpSummary struct {
	Start        int64                          `json:"start" bson:"start"`
	WarmUpDone   int64                          `json:"warmUpDone" bson:"warmUpDone"`
	End          int64                          `json:"end" bson:"end"`
	Count        int                            `json:"count" bson:"count"`
	Failed       int                            `json:"failed" bson:"failed"`
	ByPercentile []*HttpSummaryByPercentileData `json:"byPercentile" bson:"byPercentile"`
	ByMethod     []*HttpSummaryByMethodData     `json:"byMethod" bson:"byMethod"`
	ByStatus     []*HttpSummaryByStatusData     `json:"byStatus" bson:"byStatus"`
}

type HttpSummaryByPercentileData struct {
	Percentile int `json:"percentile" bson:"percentile"`
	Duration   int `json:"duration" bson:"duration"`
}

type HttpSummaryByStatusData struct {
	Status string `json:"status" bson:"status"`
	Count  int    `json:"count" bson:"count"`
}

type HttpSummaryByMethodData struct {
	Method string `json:"method" bson:"method"`
	Count  int    `json:"count" bson:"count"`
	Failed int    `json:"failed" bson:"failed"`
}

type HttpUrl struct {
	Url     string  `json:"url" bson:"url"`
	Count   int     `json:"count" bson:"count"`
	Failure float64 `json:"failure" bson:"failure"`
	Avg     int     `json:"avg" bson:"avg"`
	CV      float64 `json:"cv" bson:"cv"`
}

func NewHttpUrlPagingDataType() *PagingDataType {
	return NewPagingDataType(reflect.TypeOf(&HttpUrl{}))
}
