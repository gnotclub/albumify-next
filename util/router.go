package util

import (
    "github.com/gorilla/mux"
)

var Router *mux.Router
func GetRouter() {
    Router = mux.NewRouter()
    Router.StrictSlash(true)
}
