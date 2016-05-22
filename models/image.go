package models

type Image struct {
    Title string `bson:"title" json:"title"`
    Description string `bson:"description" json:"description"`
    Link string `bson:"link" json:"link"`
}
