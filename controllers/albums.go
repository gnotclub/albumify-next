package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "github.com/gnotclub/albumify-next/util"
    "github.com/gnotclub/albumify-next/models"
)

const prefix string = "/albums"
var collection *mgo.Collection
func AlbumRegisterController() {
    collection = util.GetDB().C("albums")

    util.Router.HandleFunc(prefix + "/{albumId}", AlbumShow)
}

func AlbumShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    albumId := vars["albumId"]
    documentId := util.UrlDecode(albumId)
    result := models.Album{}
    err := collection.Find(bson.M{"_id": documentId}).One(&result)
    if err != nil {
        http.Error(w, "not found", http.StatusNotFound)
    } else {
        json.NewEncoder(w).Encode(result)
    }
}
