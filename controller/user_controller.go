package controller




// GET a user by its ID
func AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user_dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}
