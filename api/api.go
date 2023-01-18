package api

import (
	"apiDesign/handler"
	"apiDesign/middleware"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func CreateRouter() {
	r := chi.NewRouter()

	r.Get("/", handler.Welcome)
	r.Post("/registerUser", handler.Register)
	r.Post("/logIn", handler.LogIn)

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
