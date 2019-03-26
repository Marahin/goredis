package rdb

import (
	logger "github.com/sirupsen/logrus"
)

type DataStore struct {
	data map[string]interface{}
}

func (ds *DataStore) Set(k string, v interface{}) {
	if len(ds.data) == 0 {
		ds.data = make(map[string]interface{})
	}

	ds.data[k] = v
}

func (ds *DataStore) Get(k string) interface{} {
	logger.WithFields(logger.Fields{"k": k}).Debug("rdb.Get starting")

	resp := ds.data[k]

	logger.WithFields(logger.Fields{"resp": resp}).Debug("rdb.Get returning")

	return resp
}

var GeneralDataStore = DataStore{}
