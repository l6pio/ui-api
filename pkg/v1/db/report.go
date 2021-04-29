package db

import (
	"gopkg.in/mgo.v2/bson"
	"l6p.io/ui-api/pkg/cfg"
	"l6p.io/ui-api/pkg/v1/vo"
)

func FindHttpTimeline(conf *cfg.Config, name string) ([]*vo.HttpTimeline, error) {
	session, col, err := GetCol(conf, name)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var ret []*vo.HttpTimeline
	err = col.Find(bson.M{"dataType": 1}).All(&ret)
	return ret, err
}

func FindHttpSummary(conf *cfg.Config, name string) (*vo.HttpSummary, error) {
	session, col, err := GetCol(conf, name)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var ret *vo.HttpSummary
	err = col.Find(bson.M{"dataType": 2}).One(&ret)
	return ret, err
}

func FindHttpUrls(conf *cfg.Config, name string, page int, rowsPerPage int, order string) (*vo.PagingData, error) {
	session, col, err := GetCol(conf, name)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return vo.NewPaging(col,
		&vo.PagingArgs{
			AddFields: bson.M{
				"failure": bson.M{"$divide": []string{"$failed", "$count"}},
			},
			Match:       bson.M{"dataType": 3},
			Sort:        ToSort(order),
			Page:        page,
			RowsPerPage: rowsPerPage,
			DataType:    vo.NewHttpUrlPagingDataType(),
		},
	)
}
