package db

import (
	"l6p.io/ui-api/pkg/cfg"
	"l6p.io/ui-api/pkg/v1/vo"
	"strings"
)

func GetJobs(conf *cfg.Config) (*vo.PagingData, error) {
	session, db, err := GetDB(conf)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	colNames, err := db.CollectionNames()
	if colNames == nil {
		colNames = []string{}
	}

	return &vo.PagingData{
		Slice: colNames,
		Count: len(colNames),
	}, nil
}

func FindJobsByKeyword(conf *cfg.Config, keyword string) (*vo.PagingData, error) {
	session, db, err := GetDB(conf)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	colNames, err := db.CollectionNames()

	slice := make([]string, 0)
	for _, name := range colNames {
		if strings.Contains(name, keyword) {
			slice = append(slice, name)
		}
	}

	return &vo.PagingData{
		Slice: slice,
		Count: len(slice),
	}, nil
}
