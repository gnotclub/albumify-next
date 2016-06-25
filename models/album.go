package models

import (
	"github.com/gnotclub/albumify-next/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var AlbumCollection string = "albums"

type Image struct {
	Title       string `bson:"title" json:"title" schema:"title"`
	Description string `bson:"description" json:"description" schema:"description"`
	Link        string `bson:"link" json:"link" schema:"link"`
}

// An album contains images and metadata
type Album struct {
	Id          int64   `bson:"_id" json:"_id" schema:"-"`
	Title       string  `bson:"title" json:"title" schema:"album_title_input"`
	Description string  `bson:"description" json:"description" schema:"album_desc_input"`
	Images      []Image `bson:"images" json:"images" schema:"frame"`
}

func GetAlbum(query bson.M) (error, Album) {
	var result Album
	var c *mgo.Collection = util.GetDB().C(AlbumCollection)
	err := c.Find(query).One(&result)

	return err, result
}

func PutAlbum(album *Album) error {
	var prevAlbum Album
	var c *mgo.Collection = util.GetDB().C(AlbumCollection)

	err := c.Find(bson.M{}).Sort("-_id").One(&prevAlbum)
	if err != nil {
		if err.Error() == "not found" {
			prevAlbum.Id = 0
		} else {
			return err
		}
	}

	(*album).Id = prevAlbum.Id + 1
	err = c.Insert(*album)
	if err != nil {
		return err
	}

	return nil
}
