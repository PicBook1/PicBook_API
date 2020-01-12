package models

import (
	"crypto/sha256"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)


type Gallery struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserName    string        `bson:"username" json:"username"`
	Title       string        `bson:"title" json:"title"`                  
	Discription string        `bson:"discription" json:"discription"`
	Path        string        `bson:"path" json:"path"`
	Comment     []Comment     `bson:"comment" json:"comment"`
	Rating      int
}

type Fav struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	UserId    string        `bson:"userid" json:"userid"`
	GalleryId string        `bson:"galleryid" json:"galleryid"`
}
