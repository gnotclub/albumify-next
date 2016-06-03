package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

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

	util.Logger.Printf("Listening on port %d", util.Config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(util.Config.ServerPort), handlers.LoggingHandler(os.Stdout, util.Router)))
}
