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

func FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := user_dao.FindByEmail(params["email"])

	fmt.Println(user.Email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, " You have User account")
		return
	}
	respondWithJson(w, http.StatusOK, user)
	return
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

// PUT update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := user_dao.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}



// DELETE an existing user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := user_dao.Delete(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	user_dao.Server = config.Server
	user_dao.Database = config.Database
	user_dao.Connect()
}
