package handler

import (
	"apiDesign/db"
	"apiDesign/model"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("error : %s\n", err.Error())
		return
	}
	if _, ok := db.BookList[book.UUID]; ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("book with the id already exists"))
		return
	}
	_, err = w.Write([]byte("book successfully added. "))
	db.BookList[book.UUID] = book
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Println("failed to write data in response body")
		_, err = w.Write([]byte(err.Error()))
	}
	return
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	temp := chi.URLParam(r, "bookId")
	bookId, _ := strconv.Atoi(temp)

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("failed to Decode Data From Request")
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	if _, ok := db.BookList[bookId]; ok {
		db.BookList[bookId] = book
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("failed to write data in response body")
			_, err2 := w.Write([]byte(err.Error()))
			if err2 != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Book Update Successfully\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write([]byte("bookServer have no Book with this bookID"))
	if err != nil {
		return
	}

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	temp := chi.URLParam(r, "bookId")
	bookId, _ := strconv.Atoi(temp)
	if _, ok := db.BookList[bookId]; ok {
		var book model.Book
		book = db.BookList[bookId]
		delete(db.BookList, bookId)
		_, err := w.Write([]byte("Book Delete Successfully\n"))
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			log.Println("failed to write data in response body")
			_, err = w.Write([]byte(err.Error()))
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("bookServer have no Book with this bookID"))
	if err != nil {
		return
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte("failed to Decode Data"))
		if err2 != nil {
			fmt.Printf("error : %s\n", err2.Error())
			return
		}
		return
	}
	db.UserList = append(db.UserList, user)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("failed to parse data in response body")
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			fmt.Printf("error : %s\n", err.Error())
			return
		}
		return
	}
}
