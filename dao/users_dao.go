package dao


import (
	"log"

	. "github.com/picbook1/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Find list of users
func (m *UsersDAO) FindAll() ([]User, error) {
	var users []User
	err := userdb.C(user_COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}
