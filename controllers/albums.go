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
	collection = util.GetDB().C(models.AlbumCollection)

	subrouter.HandleFunc("/{albumId}", AlbumShow).Methods("GET")
	subrouter.HandleFunc("/", AlbumSubmit).Methods("GET", "POST")
}

// Display an Album in JSON form
func AlbumShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	albumId := vars["albumId"]
	documentId := util.UrlDecode(albumId)
	err, result := models.GetAlbum(bson.M{"_id": documentId})
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

// Register a new album based on the JSON passed in the request
func AlbumSubmit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var album models.Album
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&album)
	// TODO: validation
	if err != nil {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		var s []byte
		r.Body.Read(s)
		util.Logger.Printf("%s Bad Request in album submission: %s", err)
		return
	}
	err = models.PutAlbum(&album)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		util.Logger.Printf("Error while trying to insert album into collection: %s", err)
	}
	outData, _ := json.Marshal(
		map[string]interface{}{
			"code": util.UrlEncode(album.Id),
			"url": fmt.Sprintf("%s:%d%s%s/%s", util.Config.ServerHostname, util.Config.ServerPort,
				util.Config.ApiEndpoint, prefix, util.UrlEncode(album.Id)),
		},
	)
	fmt.Fprintf(w, "%s", outData)
}
