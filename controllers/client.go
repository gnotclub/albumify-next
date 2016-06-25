package controllers

import (
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gnotclub/albumify-next/models"
	"github.com/gnotclub/albumify-next/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var IndexTemplate, ViewTemplate *template.Template

func ClientRegisterController() {
	util.MainRouter.HandleFunc("/", CreateAlbum).Methods("GET", "POST")
	util.MainRouter.HandleFunc("/{albumId}", ViewAlbum).Methods("GET")
	util.MainRouter.HandleFunc(util.Config.AssetsFolder+"/{path:.*}", ServeAssets)
}

func CompileClientTemplates() {
	IndexTemplate, _ = template.ParseFiles("views/header.html", "views/index.html", "views/footer.html")
	ViewTemplate, _ = template.ParseFiles("views/header.html", "views/view.html", "views/footer.html")
}

func ServeAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		IndexTemplate.ExecuteTemplate(w, "index", nil)
	} else {
		err := r.ParseForm()

		if err != nil {
			util.Logger.Fatal("Error while parsing form: " + err.Error())
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		var album models.Album
		decoder := schema.NewDecoder()
		err = decoder.Decode(&album, r.PostForm)
		if err != nil {
			util.Logger.Fatal("Error while parsing form: " + err.Error())
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}
		err = models.PutAlbum(&album)
		if err != nil {
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			util.Logger.Printf("Error while trying to insert album into collection: %s", err)
		}
		id := util.UrlEncode(album.Id)
		http.Redirect(w, r, "/"+id, 301)
	}
}

func ViewAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumId := vars["albumId"]
	documentId := util.UrlDecode(albumId)
	err, result := models.GetAlbum(bson.M{"_id": documentId})
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ViewTemplate.ExecuteTemplate(w, "view", result)
}
