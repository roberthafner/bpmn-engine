package database

import (
	"errors"
)

type queryMapper struct {
	queryMap map[string]string
}

func newQueryMapper() queryMapper {
	qm := queryMapper{}
	qm.queryMap = make(map[string]string)
	qm.queryMap["Deployment.insert"] =
		"insert into BR_DEPLOYMENT (DEPLOYMENT_ID, DEPLOYMENT_NAME, DEPLOYMENT_TIME) values (? ,?, ?)"
	qm.queryMap["Deployment.update"] =
		"update BR_DEPLOYMENT set DEPLOYMENT_NAME=?, DEPLOYMENT_TIME=? where DEPLOYMENT_ID=?"
	qm.queryMap["Deployment.delete"] =
		"delete from BR_DEPLOYMENT where DEPLOYMENT_ID=?"

	return qm
}

func (qm queryMapper) getStatement(key string) string {
	statement, found := qm.queryMap[key]
	if found {
		return statement
	}
	err := errors.New("No statement found for key: " + key)
	panic(err)
}
