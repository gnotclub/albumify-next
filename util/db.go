package util

import (
    "gopkg.in/mgo.v2"
)

var DBSession *mgo.Session
func GetDBSession() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    Logger.Println("Database connection openned.")
    DBSession = session
}

func GetDB() *mgo.Database {
    return DBSession.DB("albumify")
}
