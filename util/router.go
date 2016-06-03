package util

import (
	"github.com/gorilla/mux"
)

var MainRouter *mux.Router
var Router *mux.Router

func GetRouter() {
	MainRouter = mux.NewRouter()
	MainRouter.StrictSlash(true)
	Router = MainRouter.PathPrefix(Config.ApiEndpoint).Subrouter()
}
