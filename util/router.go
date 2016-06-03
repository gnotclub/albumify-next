package util

import "github.com/gorilla/mux"

var MainRouter *mux.Router
var Router *mux.Router

// Returns a new pre-configured router instance
func GetRouter() {
	MainRouter = mux.NewRouter()
	MainRouter.StrictSlash(true)
	Router = MainRouter.PathPrefix(Config.ApiEndpoint).Subrouter()
}
