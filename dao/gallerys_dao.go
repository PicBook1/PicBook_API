package dao

import (
	"log"

	. "github.com/picbook1/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type GallerysDAO struct {
	Server   string
	Database string
}

var gallerydb *mgo.Database

const (
	GLy_COLLECTION = "gallerys"
)

// Establish a connection to database
func (m *GallerysDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	gallerydb = session.DB(m.Database)
}

// Find list of gallerys
func (m *GallerysDAO) FindAll() ([]Gallery, error) {
	var gallerys []Gallery
	err := gallerydb.C(GLy_COLLECTION).Find(bson.M{}).All(&gallerys)
	return gallerys, err
}

// Find a gallery by its id
func (m *GallerysDAO) FindById(id string) (Gallery, error) {
	var gallery Gallery
	err := gallerydb.C(GLy_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&gallery)
	return gallery, err
}
func (m *GallerysDAO) FindByTitle(title string) (Gallery, error) {
	var gallery Gallery
	err := gallerydb.C(GLy_COLLECTION).Find(title).One(&gallery)
	return gallery, err
}
