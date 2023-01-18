package api

import (
	"apiDesign/handler"
	"apiDesign/middleware"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateRouter() {
	r := chi.NewRouter()

	r.Get("/", handler.Welcome)
	// Any user can register using user information, see user model on model section.
	r.Post("/registerUser", handler.Register)
	// User can logged-in using basic auth userName and password
	r.Post("/logIn", handler.LogIn)

	// First User need to log-in and generete JWT token
	// user can acces this groups function using baerier auth with JWT token
	r.Group(func(r chi.Router) {
		r.Use(middleware.VerifyJWT)
		r.Get("/books", handler.AllBookList)
		r.Post("/books", handler.AddBook)
		r.Get("/books/{bookId}", handler.FindBook)
		r.Put("/books/{bookId}", handler.UpdateBook)
		r.Delete("/books/{bookId}", handler.DeleteBook)
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("error : %s\n", err.Error())
	}
}
