package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"l6p.io/ui-api/pkg/cfg"
	"strings"
	"time"
)

func GetDB(conf *cfg.Config) (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.DialWithInfo(dialInfo(conf.DbAddr, conf.DbName, conf.DbUser, conf.DbPass))
	if err != nil {
		return nil, nil, err
	}

	db := session.DB(conf.DbName)
	return session, db, nil
}

func GetCol(conf *cfg.Config, name string) (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.DialWithInfo(dialInfo(conf.DbAddr, conf.DbName, conf.DbUser, conf.DbPass))
	if err != nil {
		return nil, nil, err
	}

	col := session.DB(conf.DbName).C(name)
	return session, col, nil
}

func ToSort(order string) bson.M {
	if strings.HasPrefix(order, "-") {
		return bson.M{order[1:]: -1}
	} else {
		return bson.M{order: 1}
	}
}

func dialInfo(dbAddr string, dbName string, username string, password string) *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:    []string{dbAddr},
		Direct:   false,
		Timeout:  time.Second * 15,
		Database: dbName,
		Source:   "admin",
		Username: username,
		Password: password,
	}
}
