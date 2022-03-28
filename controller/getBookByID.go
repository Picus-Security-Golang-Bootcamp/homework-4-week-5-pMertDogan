package controllers

import (
	
	"log"
	"net/http"

	"github.com/pMertDogan/picusGoBackend--Patika/picusWeek5/domain/book"
)

//curl -v "localhost:8080/book?bookID=1"
//get book by id
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookdID := r.URL.Query().Get("bookID")
	authorInfo := r.URL.Query().Get("authorInfo")
	log.Println("bookID: " + bookdID)
	log.Println("authorInfo: " + authorInfo)

	b, err := book.Repo().GetByID(bookdID)
	// b := book.Repo().GetByIdWithAuthorName(bookdID)
	// fmt.Println(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book is not exist"))
		//Good idea? exit?
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	v, err := b.Marshal()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Fatal(err)

	}

	w.Write([]byte(v))
}
