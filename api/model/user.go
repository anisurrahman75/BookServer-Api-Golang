package model

import "gorm.io/gorm"

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
