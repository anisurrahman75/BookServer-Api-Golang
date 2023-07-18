package db

import (
	model2 "github.com/anisurahman75/apiDesign/api/model"
	"gorm.io/gorm"
	"log"
)

var books = []model2.Book{
	{UUID: 1, Name: "golang in action", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5951"},
	{UUID: 2, Name: "programming in k8s", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5952"},
	{UUID: 3, Name: "hands on kubernetes", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5953"},
	{UUID: 4, Name: "web application in go", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5954"},
}
var users = []model2.User{
	{FirstName: "anisur", LastName: "Rahman", UserName: "sunny", Password: "123"},
	{FirstName: "Mridul", LastName: "Halder", UserName: "mridul12", Password: "123"},
}

func Load(db *gorm.DB) {
	if err := db.Migrator().DropTable(&model2.Book{}, &model2.User{}); err != nil {
		log.Fatal("Unable to drop tables. error: %w", err)
	}
	if err := db.AutoMigrate(&model2.Book{}, &model2.User{}); err != nil {
		log.Fatal("Unable to migrate tables. error: %w", err)
	}
	for _, book := range books {
		db.Create(book)
	}
	for _, user := range users {
		db.Create(user)
	}

}
