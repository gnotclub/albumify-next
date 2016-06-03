package util

import (
	"gopkg.in/mgo.v2"
)

var DBSession *mgo.Session

func GetDBSession() {
	session, err := mgo.Dial(Config.DatabaseAddress)
	if err != nil {
		Logger.Fatalf("Error while openning database connection: %s", err)
	}
	Logger.Printf("Database connection to %s openned.", Config.DatabaseAddress)
	DBSession = session
}

func GetDB() *mgo.Database {
	return DBSession.DB("albumify")
}
