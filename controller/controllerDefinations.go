package controllers

import (
	// "encoding/json"
	"log"
	"net/http"

	bookRest "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/book"
	// "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain"
	// "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//setup http controllers
func SetupControllers() {
	//curl -v "http://localhost:8080/statusCheck"
	http.HandleFunc("/statusCheck", StatusCheck)
	//curl -v "http://localhost:8080/book?bookID=1"
	http.HandleFunc("/book", bookHandler)
	//curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
	//Maybe we can change it to /book and change the handling method use path variable like /book?author=true
	http.HandleFunc("/bookWithAuthor", bookRest.GetBookByIdIncludeAuthor)
	//curl -v "http://localhost:8080/book/find?name=Hob"
	http.HandleFunc("/book/find", bookRest.FindBookByNameWithoutAuthor)
	//curl -v "http://localhost:8080/book" POST
	http.HandleFunc("/book/", bookRest.FindBookByNameWithoutAuthor)

	// http.HandleFunc("/book", bookRest.UpdateBookWithID) NOT IMPLEMENTED YET
	// log.Println(http.ListenAndServeTLS(":8080", "certFile", "keyFile", nil))

	http.HandleFunc("/author/", bookRest.FindBookByNameWithoutAuthor)


	log.Println(http.ListenAndServe(":8080", nil))
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	// var responseModel domain.APIStruct = domain.APIStruct{}
	log.Println(r.Method)
	w.Header().Set("Content-Type", "application/json")



	//redirect to method
	switch r.Method {
	case http.MethodGet:
		log.Println("get book by id")
		bookRest.GetBookByID(w, r)
		return
	case http.MethodPost:
		

		bookRest.UpdateBookQuantity(w,r)
		return
	default:
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

}
