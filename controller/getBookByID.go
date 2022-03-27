package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//curl -v "localhost:8080/book?bookID=1"
//get book by id
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookdID := r.URL.Query().Get("bookID")
	log.Println("bookID: " + bookdID)

	b, err := book.Singleton().GetByID(bookdID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book is not exist"))
		//Good idea? exit?
		log.Fatal(err)
	}
	
	fmt.Println(b)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(b.String()))
}
