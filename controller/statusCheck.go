package controllers

import "net/http"

//curl -v "http://localhost:8080/statusCheck"
//To check our service is running 
func StatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is running"))
}