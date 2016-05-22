package main

import (
    "fmt"
    "log"
    "os"
    "net/http"

    "github.com/gorilla/handlers"

    "github.com/tudurom/albumify-next/util"
    "github.com/tudurom/albumify-next/controllers"
)

func main() {
    util.GetLogger()

    util.GetDBSession()
    defer util.DBSession.Close()

    util.GetRouter()
    util.Router.HandleFunc("/", EmptyHandler)

    controllers.AlbumRegisterController()

    util.Logger.Println("Server ready.")
    log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, util.Router)))
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "")
}
