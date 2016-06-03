package util

import (
	"net/http"

	"github.com/gorilla/mux"
)

var MainRouter *mux.Router
var Router *mux.Router

// Returns a new pre-configured router instance
func GetRouter() {
	MainRouter = mux.NewRouter()
	MainRouter.StrictSlash(true)
	Router = MainRouter.PathPrefix(Config.ApiEndpoint).Subrouter()
}

type MyServer struct {
	R *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.R.ServeHTTP(rw, req)
}
