package controllers

import "net/http"

//curl -v "localhost:8080/book?bookID=1"
//get book by id
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("bookID")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param))
}
