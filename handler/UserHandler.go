package handler

import (
	"apiDesign/db"
	"apiDesign/middleware"
	"apiDesign/model"
	"encoding/json"
	"fmt"

	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello\nWelcome to My Book-Server!!!"))
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Welcome End-Point")
		return
	}
	return
}
func AllBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(db.BookList)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From AllBookList End-Point")
		return
	}
}
func FindBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	temp := chi.URLParam(r, "bookId")
	bookId, _ := strconv.Atoi(temp)
	if _, ok := db.BookList[bookId]; ok {
		var book model.Book
		book = db.BookList[bookId]
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Book Find Successfully\n"))
		if err != nil {
			return
		}
		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			model.RequestForError(http.StatusBadRequest, err.Error(), w, "From FindBook End-Point: Failed to write data in response body")
			return
		}
	}
	model.RequestForError(http.StatusBadRequest, "", w, "From FindBook End-Point: BookServer have no Book with this bookID")
}
func checkValidUser(userName, password string) bool {
	for _, i := range db.UserList {
		if userName == i.UserName && password == i.Password {
			return true
		}
	}
	return false
}
func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if ok {
		if !checkValidUser(username, password) {
			model.RequestForError(http.StatusUnauthorized, "", w, "From LogIn End-Point: Invalid UserName or Password")
			return
		}
		str, _ := middleware.GenerateJWT(username)
		// set token on cookies storage
		http.SetCookie(w, &http.Cookie{Name: "jwt", Value: str, Expires: time.Now().Add(10 * time.Minute)})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint("LogIn Successully\nBearer Token Below\n", str)))
		return
	}
	model.RequestForError(http.StatusUnauthorized, "", w, "From LogIn End-Point: Invalid UserName or Password")
}
