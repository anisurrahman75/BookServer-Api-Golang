package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type Book struct {
	UUID        int    `gorm:"primary_key" json:"uuid"`
	Name        string `gorm:"size:255" json:"name"`
	Author      string `gorm:"size:255;not null;"json:"author"`
	PublishDate string `gorm:"size:255" json:"publishDate"`
	ISBN        string `gorm:"size:255; unique" json:"ISBN"`
}

func (b *Book) Exist(db *gorm.DB) bool {
	if result := db.First(&b, b.UUID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		log.Fatal(result.Error)
	}
	return true
}

func (b *Book) Create(db *gorm.DB) int64 {
	result := db.Create(&b)
	if result.Error != nil {
		log.Fatalf("%w", result.Error)
	}
	return result.RowsAffected
}

func (b *Book) Update(db *gorm.DB) int64 {

	result := db.Model(Book{}).Where("uuid = ?", b.UUID).Updates(&b)
	if result.Error != nil {
		log.Fatalf("%w", result.Error)
	}
	return result.RowsAffected
}

func (b *Book) Delete(db *gorm.DB) int64 {
	result := db.Delete(&b, b.UUID)
	if result.Error != nil {
		log.Fatalf("%w", result.Error)
	}
	return result.RowsAffected
}

func (b *Book) AllBooks(db *gorm.DB) (*[]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return &books, nil
}
