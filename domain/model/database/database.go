package database

import (
	"github.com/roberthafner/bpmn-engine/domain/model"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
	"strings"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

func NewDatabase() Database {
	dbs := Database{}
	var err error

	var settings = sqlite.ConnectionURL{
		Database: "./burrow.db",
	}

	dbs.db, err = sqlite.Open(settings)
	if err != nil {
		panic(err)
	}
	dbs.db.SetLogging(true)

	fmt.Println(os.Getwd())
	// Initialize schema
	file, err := ioutil.ReadFile("../model/database/schema.sql")
	if err != nil {
		panic(err)
	}

	statements := strings.Split(string(file), ";")

	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if len(statement) > 0 {
			_, err := dbs.db.Exec(statement)
			if err != nil {
				panic(err)
			}
		}
	}

	//dbs.qm = newQueryMapper()
	return dbs
}

type Database struct {
	db sqlbuilder.Database
}

func (db Database) Insert(e model.Entity) {
	_, err := db.db.Collection(getCollectionName(e)).Insert(e)
	if err != nil {
		panic(err)
	}
}

func (db Database) Update(e model.Entity) {
	err := db.db.Collection(getCollectionName(e)).UpdateReturning(&e)
	if err != nil {
		panic(err)
	}
}

func (dbs Database) Delete(e model.Entity) {
	dbs.db.Collection(getCollectionName(e)).Find().Delete()
}

func (dbs Database) Close() {
	err := dbs.db.Close()
	if err != nil {
		panic(err)
	}
}

func getCollectionName(e model.Entity) string {
	switch t := e.(type) {
	case model.DeploymentEntity:
		return "BR_DEPLOYMENT"
	case model.DeploymentResource:
		return "BR_DEPLOYMENT_RESOURCE"
	default:
		panic(errors.New(fmt.Sprintf("Unsupported entity type: %s", t)))
	}
}
