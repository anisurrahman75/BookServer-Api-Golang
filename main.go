package main

import (
	"apiDesign/api"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func main() {
	// create router by calling CreateRouter function
	s := api.CreateNewServer()
	s.MountHandlers()
	err := http.ListenAndServe(":3000", s.Router)
	if err != nil {
		fmt.Printf("error : %s\n", err.Error())
	}
}
