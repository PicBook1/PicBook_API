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
