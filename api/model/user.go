package model

import "gorm.io/gorm"

type User struct {
	FirstName string `gorm:"size:255" json:"firstName"`
	LastName  string `gorm:"size:255" json:"lastName"`
	UserName  string `gorm:"primary_key; unique" json:"userName"`
	Password  string `gorm:"size:255" json:"password"`
}

func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
