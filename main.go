package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/picbook1/controller"


)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/users", controller.AllUsers).Methods("GET")
	r.HandleFunc("/users", controller.UpdateUser).Methods("PUT")
	r.HandlerFunc("/users", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", controller.FindUser).Methods("GET")
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	
if err := http.ListenAndServe(":8008", r); err != nil {
		log.Fatal(err)
	}

}






