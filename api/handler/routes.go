package handler

import (
	"github.com/anisurahman75/apiDesign/api/db"
	"github.com/anisurahman75/apiDesign/api/middleware"
	"github.com/go-chi/chi/v5"
	mi "github.com/go-chi/chi/v5/middleware"
)

var Port string
var Auth bool

func (s *Server) MountHandlers() {

	db.UserInit()
	db.BookInit()
	db.Load(s.DB)

	s.Router.Use(mi.Logger)
	s.Router.Get("/api", s.Welcome)
	//Any user can register using user information, see user model on model section.
	s.Router.Post("/api/registerUser", s.Register)
	// User can logged-in using basic auth userName and password
	s.Router.Post("/api/logIn", s.LogIn)
	// First User need to log-in and generete JWT token
	// user can acces this groups function using berier auth with JWT token only
	s.Router.Group(func(r chi.Router) {
		if Auth == true {
			r.Use(middleware.VerifyJWT)
		}
		r.Get("/api/books", s.BookList)
		r.Post("/api/books", s.AddBook)
		r.Get("/api/books/{bookId}", s.FindBook)
		r.Put("/api/books/{bookId}", s.UpdateBook)
		r.Delete("/api/books/{bookId}", s.DeleteBook)
	})
}
