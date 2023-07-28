package model

import "gorm.io/gorm"

type Book struct {
	UUID        int    `gorm:"primary_key" json:"uuid"`
	Name        string `gorm:"size:255" json:"name"`
	Author      string `gorm:"size:255;not null;"json:"author"`
	PublishDate string `gorm:"size:255" json:"publishDate"`
	ISBN        string `gorm:"size:255; unique" json:"ISBN"`
}

func (b *Book) Exist() {
	
}

func (b *Book) GetAllBooks(db *gorm.DB) (*[]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return &books, nil
}
