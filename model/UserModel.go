package model

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	UserName string `json:"userName"`
	Password string `json:"password"`
}
