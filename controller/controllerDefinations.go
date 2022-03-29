package controllers

import (
	bookRest "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/book"
	"log"
	"net/http"
)

//setup http controllers
func SetupControllers() {
	//curl -v "http://localhost:8080/ready"
	http.HandleFunc("/statusCheck", StatusCheck)
	//curl -v "http://localhost:8080/book?bookID=1"
	http.HandleFunc("/book", bookRest.GetBookByID)
	//curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
	//Maybe we can change it to /book and change the handling method use path variable like /book?author=true
	http.HandleFunc("/bookWithAuthor", bookRest.GetBookByIdIncludeAuthor)
	//curl -v "http://localhost:8080/book/find?name=Hob"
	http.HandleFunc("/book/find", bookRest.FindBookByNameWithoutAuthor)

	
	// http.HandleFunc("/book", bookRest.UpdateBookWithID) NOT IMPLEMENTED YET
	// log.Println(http.ListenAndServeTLS(":8080", "certFile", "keyFile", nil))

	log.Println(http.ListenAndServe(":8080", nil))
}
