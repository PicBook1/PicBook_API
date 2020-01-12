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


// Update an existing user
func (m *UsersDAO) Update(user User) error {
	err := userdb.C(user_COLLECTION).UpdateId(user.Name, &user)
	return err
}

// Find a user by its name

func (m *UsersDAO) FindByName(name string) (User, error) {
	var user User
	err := userdb.C(user_COLLECTION).Find(name).One(&user)
	return user, err
}
