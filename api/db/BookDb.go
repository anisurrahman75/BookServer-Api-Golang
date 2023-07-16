package db

import (
	"github.com/anisurahman75/apiDesign/api/model"
)

var BookList = map[int]model.Book{}

func BookInit() {
	BookList = map[int]model.Book{
		1: {UUID: 1, Name: "learn-api", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5954"},
		2: {UUID: 2, Name: "learn-api", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5954"},
		3: {UUID: 3, Name: "learn-api", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5954"},
		4: {UUID: 4, Name: "learn-api", Author: "Anisur", PublishDate: "01-02-2022", ISBN: "0999-0555-5954"},
	}
}
