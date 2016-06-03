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

	var configFile = flag.String("config", "config.json", "path of the config file")
	flag.Parse()
	util.ReadConfig(*configFile)

	util.GetDBSession()
	defer util.DBSession.Close()

	util.GetRouter()

	controllers.AlbumRegisterController()

	address := fmt.Sprintf("%s:%d", util.Config.ServerHostname, util.Config.ServerPort)
	util.Logger.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, util.Router)))
}
