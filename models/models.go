package models

import (
	"crypto/sha256"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Gallery  []Gallery     `bson:"gallery" json:"gallery"`
	Fav      []Fav         `bson:"fav" json:"fav"`
}

type Gallery struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserName    string        `bson:"username" json:"username"`
	Title       string        `bson:"title" json:"title"`                  
	Discription string        `bson:"discription" json:"discription"`
	Path        string        `bson:"path" json:"path"`
	Comment     []Comment     `bson:"comment" json:"comment"`
	Rating      int
}

type Comment struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	GalleryId string        `bson:"galleryid" json:"galleryid"`            
	Name      string        `bson:"name" json:"name"`
	Message   string        `bson:"message" json:"message"`
}

type Fav struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	UserId    string        `bson:"userid" json:"userid"`
	GalleryId string        `bson:"galleryid" json:"galleryid"`
}
