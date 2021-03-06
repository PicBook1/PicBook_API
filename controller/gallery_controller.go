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


// GET a gallery by its ID
func AllGallerys(w http.ResponseWriter, r *http.Request) {
	gallerys, err := gallery_dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, gallerys)
}



// GET a gallery by its ID
func FindGallery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gallery, err := gallery_dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Gallery ID")
		return
	}
	respondWithJson(w, http.StatusOK, gallery)
}

func FindGalleryBytitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gallery, err := gallery_dao.FindByTitle(params["title"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, " You have Gallery account")
		return
	}
	respondWithJson(w, http.StatusOK, gallery)
}
// PUT update an existing gallery
func UpdateGallery(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gallery Gallery
	if err := json.NewDecoder(r.Body).Decode(&gallery); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := gallery_dao.Update(gallery); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
// POST a new gallery
func CreateGallery(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gallery Gallery
	if err := json.NewDecoder(r.Body).Decode(&gallery); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	gallery.ID = bson.NewObjectId()
	if err := gallery_dao.Insert(gallery); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err:= FindGalleryBytitle; err != nil {
		respondWithJson(w, http.StatusCreated, gallery)
		return
	}
	respondWithError(w, http.StatusBadRequest, "u have an account")

}


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


// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	gallery_dao.Server = config.Server
	gallery_dao.Database = config.Database
	gallery_dao.Connect()
}
