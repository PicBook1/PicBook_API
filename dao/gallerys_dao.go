package dao

import (
	"log"

	. "github.com/picbook1/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


// Find list of gallerys
func (m *GallerysDAO) FindAll() ([]Gallery, error) {
	var gallerys []Gallery
	err := gallerydb.C(GLy_COLLECTION).Find(bson.M{}).All(&gallerys)
	return gallerys, err
}
