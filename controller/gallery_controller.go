package controller

import (
	"encoding/json"
	. "github.com/picbook1/dao"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	. "github.com/picbook1/models"

)

//var config = Config{}
var gallery_dao = GallerysDAO{}

// DELETE an existing gallery
func DeleteGallery(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gallery Gallery
	if err := json.NewDecoder(r.Body).Decode(&gallery); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := gallery_dao.Delete(gallery); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
