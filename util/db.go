package util

import (
	"gopkg.in/mgo.v2"
)

var DBSession *mgo.Session

// Try to connect to the database
func GetDBSession() {
	session, err := mgo.Dial(Config.DatabaseAddress)
	if err != nil {
		Logger.Fatalf("Error while openning database connection: %s", err)
	}
	Logger.Printf("Database connection to %s openned. Databse name: %s", Config.DatabaseAddress, Config.DatabaseName)
	DBSession = session
}

// Returns the database we're operating on
func GetDB() *mgo.Database {
	return DBSession.DB(Config.DatabaseName)
}
