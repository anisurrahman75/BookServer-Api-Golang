package db

import "github.com/anisurahman75/apiDesign/model"

var UserList = []model.User{}

func UserInit() {
	UserList = []model.User{
		{FirstName: "anisur", LastName: "Rahman", UserName: "sunny", Password: "123"},
		{FirstName: "Mridul", LastName: "Halder", UserName: "mridul12", Password: "123"},
	}
}
