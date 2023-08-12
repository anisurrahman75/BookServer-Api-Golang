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
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, "", w, "From AddBook End-Point: Failed to write data in response body")
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("book successfully added. "))
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
	if bookId != book.UUID {
		model.RequestForError(http.StatusBadRequest, "", w, "From UpdateBook End-Point: UrlParam BookID doesn't match to Reques BookId")
		return
	}
	isUpdated := book.Update(s.DB)
	if isUpdated == int64(0) {
		model.RequestForError(http.StatusBadRequest, "", w, "From UpdateBook End-Point: Failed to Update in Database table.")
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From UpdateBook End-Point: Failed to write data in response body")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Book Update Successfully\n"))
}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var book model.Book
	var err error
	temp := chi.URLParam(r, "bookId")
	book.UUID, err = strconv.Atoi(temp)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From DeleteBook End-Point: Failed to Convert bookId")
		return
	}
	if !book.Exist(s.DB) {
		model.RequestForError(http.StatusBadRequest, "", w, "From DeleteBook End-Point: rest-api-server have no Book with this bookID")
		return
	}
	isDeleted := book.Delete(s.DB)
	if isDeleted == int64(0) {
		model.RequestForError(http.StatusBadRequest, "", w, "From DeleteBook End-Point: Failed to Delete in Database table.")
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From DeleteBook End-Point: Failed to Encode Data To Response")
	}
	_, _ = w.Write([]byte("Book Delete Successfully\n"))
	w.WriteHeader(http.StatusOK)
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

	if user.Exist(s.DB) {
		model.RequestForError(http.StatusBadRequest, "", w, "From Register End-Point: User with this UserId already exists")
		return
	}

	isCreated := user.Create(s.DB)
	if isCreated == int64(0) {
		model.RequestForError(http.StatusBadRequest, err.Error(), w, "From Register End-Point: Failed to Register User in Database table.")
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		model.RequestForError(http.StatusBadRequest, "", w, "From Register End-Point: Failed to write data in response body")
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User Register successfully. "))
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
		model.RequestForError(http.StatusBadRequest, d.Error.Error(), w, "From FindBook End-Point: rest-api-server have no Book with this bookID")
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
	books, err := book.AllBooks(s.DB)
	if err != nil {
		return books, err
	}
	return books, nil
}
