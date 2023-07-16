package model

import "gorm.io/gorm"

type Book struct {
	UUID        int    `json:"uuid"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishDate string `json:"publishDate"`
	ISBN        string `json:"ISBN"`
}

func (u *Book) GetAllBooks(db *gorm.DB) (*[]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return &books, nil
}
