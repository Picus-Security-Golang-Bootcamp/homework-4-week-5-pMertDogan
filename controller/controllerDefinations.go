package controllers

import (
	"log"
	"net/http"
)


//setup http controllers
func SetupControllers() {
	//curl -v "localhost:8080/ready"
	http.HandleFunc("/statusCheck", StatusCheck)
	//curl -v "localhost:8080/param?bookID=1"
	http.HandleFunc("/book", GetBookByID)
	http.HandleFunc("/bookWithAuthor", GetBookByIdIncludeAuthor)

	// log.Println(http.ListenAndServeTLS(":8080", "certFile", "keyFile", nil))

	log.Println(http.ListenAndServe(":8080", nil))
}
