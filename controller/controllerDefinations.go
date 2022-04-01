package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/admin"
	authorRest "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/author"
	bookRest "github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/controller/book"
)

//setup http controllers
func SetupControllers() {
	r := mux.NewRouter()

	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	//set handlers whitelist if needed
	// handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH"})

	//*******Without gorilla Mux//can be moved another file
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
	http.HandleFunc("/book/enable", bookRest.EnableBook)

	//curl -v "http://localhost:8080/statusCheck"
	r.HandleFunc("/statusCheck", StatusCheck)
	//curl -v "http://localhost:8080/book?bookID=1"
	r.HandleFunc("/book", bookHandler)
	//curl -v "http://localhost:8080/bookWithAuthor?bookID=1"
	//Maybe we can change it to /book and change the handling method use path variable like /book?author=true
	r.HandleFunc("/bookWithAuthor", bookRest.GetBookByIdIncludeAuthor)
	//curl -v "http://localhost:8080/book/find?name=Hob"
	r.HandleFunc("/book/find", bookRest.FindBookByNameWithoutAuthor)
	//curl -v "http://localhost:8080/book" POST
	r.HandleFunc("/book/enable", bookRest.EnableBook)

	// http.HandleFunc("/book", bookRest.UpdateBookWithID) NOT IMPLEMENTED YET
	//******Middleware

	// r.Use(admin.AdminLogger)
	//add admin middleware
	// r.Use(admin.AdminMiddleware)

	//**** Author Handles it with Gorilla Mux
	//setup admin routes
	authorRest.AuthorHandler(r)

	//*****setup admin routes
	admin.AdminHandler(r)

	//register gorilla mux
	http.Handle("/", r)

	// srv := &http.Server{
	// 	Addr:         "0.0.0.0:8080",
	// 	WriteTimeout: time.Second * 15,
	// 	ReadTimeout:  time.Second * 15,
	// 	IdleTimeout:  time.Second * 60,
	// 	Handler:      r,
	// }
	// log.Println(srv.ListenAndServe())

	// log.Println(http.ListenAndServeTLS(":8080", "certFile", "keyFile", nil))
	log.Println(http.ListenAndServe(":8080", r))

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
		bookdID := r.URL.Query().Get("bookID")
		if bookdID != "" {
			bookRest.UpdateBookQuantity(w, r)
		} else {
			//somethink like
			// bookRest.UpdateBook(w, r)
		}
		return

	case http.MethodDelete:
		bookRest.SoftDelete(w, r)
		return
	default:
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

}
