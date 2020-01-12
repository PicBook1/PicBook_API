package controller

import (
	"encoding/json"
	//"log"
	"net/http"

	. "github.com/picbook1/config"
	. "github.com/picbook1/dao"
	. "github.com/picbook1/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

var config = Config{}
var user_dao = UsersDAO{}


// GET a user by its ID
func AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user_dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

// GET a user by its ID
func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := user_dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}
// POST a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := user_dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithError(w, http.StatusBadRequest, "u have an account")

}
