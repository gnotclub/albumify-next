package models

// An album contains images and metadata
type Album struct {
	Id          int64    `bson:"_id" json:"_id"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	Images      []*Image `bson:"images" json:"images"`
}
