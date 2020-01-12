package dao


import (
	"log"

	. "github.com/picbook1/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var userdb *mgo.Database

const (
	user_COLLECTION = "users"
)

// Establish a connection to database
func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	userdb = session.DB(m.Database)
}

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

// Find a user by its email
func (m *UsersDAO) FindByEmail(email string) (User, error) {
	var user User
	err := userdb.C(user_COLLECTION).Find(email).One(&user)
	return user, err
}

// Insert a user into database
func (m *UsersDAO) Insert(user User) error {
	err := userdb.C(user_COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UsersDAO) Delete(user User) error {
	err := userdb.C(user_COLLECTION).Remove(&user)
	return err
}
