package api

import (
	"github.com/anisurahman75/apiDesign/db"
	"github.com/anisurahman75/apiDesign/handler"
	"github.com/anisurahman75/apiDesign/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}
func (r *Server) MountHandlers() {
	db.UserInit()
	db.BookInit()
	r.Router.Get("/api", handler.Welcome)

	//Any user can register using user information, see user model on model section.
	r.Router.Post("/api/registerUser", handler.Register)
	// User can logged-in using basic auth userName and password
	r.Router.Post("/api/logIn", handler.LogIn)

	// First User need to log-in and generete JWT token
	// user can acces this groups function using berier auth with JWT token only
	r.Router.Group(func(r chi.Router) {
		r.Use(middleware.VerifyJWT)
		r.Get("/api/books", handler.AllBookList)
		r.Post("/api/books", handler.AddBook)
		r.Get("/api/books/{bookId}", handler.FindBook)
		r.Put("/api/books/{bookId}", handler.UpdateBook)
		r.Delete("/api/books/{bookId}", handler.DeleteBook)
	})
}

//package main
//
//import (
//"fmt"
//"github.com/anisurahman75/apiDesign/api"
//"net/http"
//)
//
//func main() {
//	// create router by calling CreateRouter function
//
//}
