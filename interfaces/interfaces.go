package interfaces

import "github.com/picbook1/models"



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


