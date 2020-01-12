package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/picbook1/controller"


)

func main() {
	r := mux.NewRouter()
	
if err := http.ListenAndServe(":8008", r); err != nil {
		log.Fatal(err)
	}

}
