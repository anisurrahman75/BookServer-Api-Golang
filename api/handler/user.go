package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anisurahman75/apiDesign/api/middleware"
	"github.com/anisurahman75/apiDesign/api/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) Welcome(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello\nWelcome to My Book-Server!!!"))
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Welcome End-Point")
		return
	}
	return
}
func (s *Server) BookList(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	books, err := BookList(s)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From AllBookList End-Point")
		return
	}
	if err := json.NewEncoder(w).Encode(books); err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From AllBookList End-Point")
		return
	}
}
func (s *Server) FindBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	temp := chi.URLParam(r, "bookId")
	bookId, _ := strconv.Atoi(temp)
	var book model.Book
	d := s.DB.First(&book, "id= ?", bookId)
	if errors.Is(d.Error, gorm.ErrRecordNotFound) {
		model.RequestForError(http.StatusBadRequest, d.Error.Error(), w, "From FindBook End-Point: BookServer have no Book with this bookID")
	}
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
	return

}

func (s *Server) checkValidUser(userName, password string) (bool, error) {
	users, err := UserList(s)
	if err != nil {
		return false, err
	}
	for _, i := range *users {
		if userName == i.UserName && password == i.Password {
			return true, nil
		}
	}
	return false, nil
}

func (s *Server) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if ok {
		find, err := s.checkValidUser(username, password)
		if err != nil {
			model.RequestForError(http.StatusUnauthorized, err.Error(), w, "From LogIn End-Point.")
		}
		if !find {
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
