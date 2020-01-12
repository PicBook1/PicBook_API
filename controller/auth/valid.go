package auth

import (

	"net/http"
	"fmt"
)

func VerifyCredentials(w http.ResponseWriter, r *http.Request , name string) bool {

	fmt.Println(w,r,name)
	return false
}