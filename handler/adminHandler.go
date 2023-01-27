package handler

import (
	"encoding/json"
	"github.com/anisurahman75/apiDesign/db"
	"github.com/anisurahman75/apiDesign/model"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From AddBook End-Point: Json Data Decode Failed")
		return
	}
	if _, ok := db.BookList[book.UUID]; ok {
		model.RequestForError(http.StatusBadRequest, "", w, "From AddBook End-Point: Book with this ID already exists")
		return
	}
	_, _ = w.Write([]byte("book successfully added. "))
	db.BookList[book.UUID] = book
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, "", w, "From AddBook End-Point: Failed to write data in response body")
		return
	}
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	temp := chi.URLParam(r, "bookId")
	bookId, err := strconv.Atoi(temp)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From UpdateBook End-Point: Failed to Failed to Convert bookId")
		return
	}
	var book model.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From UpdateBook End-Point: Failed to Decode Data From Request")
		return
	}
	if _, ok := db.BookList[bookId]; ok {
		db.BookList[bookId] = book
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			model.RequestForError(http.StatusBadRequest, err.Error(), w, "From UpdateBook End-Point: Failed to Encode Data To Response")
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Book Update Successfully\n"))
		return
	}
	model.RequestForError(http.StatusBadRequest, "", w, "From UpdateBook End-Point: BookServer have no Book with this bookID")
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	temp := chi.URLParam(r, "bookId")
	bookId, err := strconv.Atoi(temp)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From DeleteBook End-Point: Failed to Convert bookId")
		return
	}
	if _, ok := db.BookList[bookId]; ok {
		var book model.Book
		book = db.BookList[bookId]
		delete(db.BookList, bookId)
		_, _ = w.Write([]byte("Book Delete Successfully\n"))
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			model.RequestForError(http.StatusBadRequest, err.Error(), w, "From DeleteBook End-Point: Failed to Encode Data To Response")
			return
		}
		return
	}
	model.RequestForError(http.StatusBadRequest, "", w, "From DeleteBook End-Point: BookServer have no Book with this bookID")
}
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Register End-Point: Json Data Decode Failed from Request")
		return
	}
	db.UserList = append(db.UserList, user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Register Successfully\n"))
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Register End-Point: Json Data Encode Failed to Response")
		return
	}
}
