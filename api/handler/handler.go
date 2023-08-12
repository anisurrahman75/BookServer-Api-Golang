package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anisurahman75/apiDesign/api/middleware"
	"github.com/anisurahman75/apiDesign/api/model"
	"github.com/anisurahman75/apiDesign/db"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.DB = db.Initialize()
	return s
}

func (s *Server) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From AddBook End-Point: Json Data Decode Failed")
		return
	}
	if book.Exist(s.DB) {
		model.RequestForError(http.StatusBadRequest, "", w, "From AddBook End-Point: Book with this ID already exists")
		return
	}
	isCreated := book.Create(s.DB)
	if isCreated == int64(0) {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From AddBook End-Point: Failed to Add book in Database table.")
		return
	}
	_, _ = w.Write([]byte("book successfully added. "))
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, "", w, "From AddBook End-Point: Failed to write data in response body")
	}

}

func (s *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {
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
	s.DB.Model(model.Book{}).Where("uuid = ?", bookId).Updates(book)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From UpdateBook End-Point: Failed to Encode Data To Response")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Book Update Successfully\n"))
}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	temp := chi.URLParam(r, "bookId")
	bookId, err := strconv.Atoi(temp)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From DeleteBook End-Point: Failed to Convert bookId")
		return
	}
	var queryBook model.Book
	queryErr := s.DB.First(&queryBook, bookId)
	if queryErr != nil && errors.Is(queryErr.Error, gorm.ErrRecordNotFound) {
		model.RequestForError(http.StatusBadRequest, "", w, "From DeleteBook End-Point: BookServer have no Book with this bookID")
		return
	}
	s.DB.Delete(&queryBook, bookId)

	_, _ = w.Write([]byte("Book Delete Successfully\n"))
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(queryBook)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From DeleteBook End-Point: Failed to Encode Data To Response")
	}
	return

}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Register End-Point: Json Data Decode Failed from Request")
		return
	}
	var queryUser model.User
	queryErr := s.DB.First(&queryUser, user.UserName)

	if err != nil {
		if errors.Is(queryErr.Error, gorm.ErrRecordNotFound) {
			s.DB.Create(&user)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Register Successfully\n"))
			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Register End-Point: Json Data Encode Failed to Response")
			}
		} else {
			model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Register End-Point.")
		}
	}
}

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
	d := s.DB.First(&book, bookId)
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

func (s *Server) CheckValidUser(userName, password string) bool {
	var result model.User
	err := s.DB.First(&result, userName)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return false
	}
	if result.Password == password {
		return true
	}
	return false
}

func (s *Server) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if ok {
		find := s.CheckValidUser(username, password)
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

func BookList(s *Server) (*[]model.Book, error) {
	book := model.Book{}
	books, err := book.GetAllBooks(s.DB)
	if err != nil {
		return books, err
	}
	return books, nil
}
