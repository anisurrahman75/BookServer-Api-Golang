package api

import (
	"apiDesign/db"
	"apiDesign/handler"
	"apiDesign/middleware"
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
	r.Router.Get("/", handler.Welcome)

	//Any user can register using user information, see user model on model section.
	r.Router.Post("/registerUser", handler.Register)
	// User can logged-in using basic auth userName and password
	r.Router.Post("/logIn", handler.LogIn)

	// First User need to log-in and generete JWT token
	// user can acces this groups function using baerier auth with JWT token only
	r.Router.Group(func(r chi.Router) {
		r.Use(middleware.VerifyJWT)
		r.Get("/books", handler.AllBookList)
		r.Post("/books", handler.AddBook)
		r.Get("/books/{bookId}", handler.FindBook)
		r.Put("/books/{bookId}", handler.UpdateBook)
		r.Delete("/books/{bookId}", handler.DeleteBook)
	})

}
