package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/gnotclub/albumify-next/controllers"
	"github.com/gnotclub/albumify-next/util"
)

func main() {
	util.GetLogger()

	// Get config file path from command line args
	var configFile = flag.String("config", "config.json", "path of the config file")
	flag.Parse()
	util.ReadConfig(*configFile)

	util.GetDBSession()
	defer util.DBSession.Close()

	util.GetRouter()

	// Register routes for all of Album's controllers
	controllers.AlbumRegisterController()
	// Register main client's controllers
	controllers.CompileClientTemplates()
	controllers.ClientRegisterController()

	// Listen
	address := fmt.Sprintf("%s:%d", util.Config.ServerHostname, util.Config.ServerPort)
	util.Logger.Printf("Listening on %s", address)
	http.Handle("/", &util.MyServer{util.MainRouter})
	log.Fatal(http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, &util.MyServer{util.MainRouter})))
}
