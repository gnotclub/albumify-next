package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gnotclub/albumify-next/models"
	"github.com/gnotclub/albumify-next/util"
)

const prefix string = "/albums"

var subrouter *mux.Router
var collection *mgo.Collection

// Registers the routes for this controller
func AlbumRegisterController() {
	subrouter = util.Router.PathPrefix(prefix).Subrouter()
	collection = util.GetDB().C("albums")

	subrouter.HandleFunc("/{albumId}", AlbumShow).Methods("GET")
	subrouter.HandleFunc("/", AlbumSubmit).Methods("GET", "POST")
}

// Display an Album in JSON form
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

// Register a new album based on the JSON passed in the request
func AlbumSubmit(w http.ResponseWriter, r *http.Request) {
	var album, prevAlbum models.Album
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&album)
	// TODO: validation
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		util.Logger.Printf("Bad Request in album submission: %s", err)
	}
	// Get the last album registered
	err = collection.Find(bson.M{}).Sort("-_id").One(&prevAlbum)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		util.Logger.Printf("Couldn't get last album in collection: %s", err)
	}
	// Increment new album's id
	album.Id = prevAlbum.Id + 1
	err = collection.Insert(album)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		util.Logger.Printf("Couldn't insert album in collection: %s", err)
	}
	fmt.Fprintf(w, "{\"code\": \"%s\"}", util.UrlEncode(album.Id))
}
