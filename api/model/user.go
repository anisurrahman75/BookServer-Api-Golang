package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type User struct {
	FirstName string `gorm:"size:255" json:"firstName"`
	LastName  string `gorm:"size:255" json:"lastName"`
	UserName  string `gorm:"primary_key; unique" json:"userName"`
	Password  string `gorm:"size:255" json:"password"`
}

func (u *User) Exist(db *gorm.DB) bool {
	if result := db.Where("user_name = ?", u.UserName).First(&u); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		log.Fatal(result.Error)
	}
	return true
}

func (u *User) Create(db *gorm.DB) int64 {
	result := db.Create(&u)
	if result.Error != nil {
		log.Fatalf("%w", result.Error)
	}
	return result.RowsAffected
}

func (u *User) AllUsers(db *gorm.DB) (*[]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
