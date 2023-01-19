package handler

import (
	"apiDesign/db"
	"apiDesign/middleware"
	"apiDesign/model"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello\nWelcome to My Book-Server!!!"))
	if err != nil {
		fmt.Printf("error : %s\n", err.Error())
	}
	return
}

func AllBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Println(r.Context().Value("user"))
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(db.BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
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
func checkValidUser(userName, password string) bool {
	for _, i := range db.UserList {
		if userName == i.UserName && password == i.Password {
			return true
		}
	}
	return false
}
func LogIn(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		if !checkValidUser(username, password) {
			http.Error(w, "Wrong UserName or Password", http.StatusUnauthorized)
			return
		}
		str, _ := middleware.GenerateJWT(username)
		// set token on cookies storage with 10 minutes time limit
		http.SetCookie(w, &http.Cookie{Name: "jwt", Value: str, Expires: time.Now().Add(10 * time.Minute)})
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Login Successfully\n"))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
		return
	}
	http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
}
