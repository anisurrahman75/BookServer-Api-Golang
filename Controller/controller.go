package Controller

import (
	"apiDesign/Method"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Createrouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", Method.Welcome)
	r.Get("/books", Method.AllBookList)
	r.Post("/books", Method.AddBook)

	http.ListenAndServe(":3000", r)
}
