package handler

import (
	"github.com/anisurahman75/apiDesign/api/db"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
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
