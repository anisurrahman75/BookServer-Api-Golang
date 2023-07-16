package handler

import (
	"github.com/anisurahman75/apiDesign/api/model"
)

func BookList(s *Server) (*[]model.Book, error) {
	book := model.Book{}
	books, err := book.GetAllBooks(s.DB)
	if err != nil {
		return books, err
	}
	return books, nil
}
func UserList(s *Server) (*[]model.User, error) {
	user := model.User{}
	users, err := user.GetAllUsers(s.DB)
	if err != nil {
		return users, err
	}
	return users, nil
}
