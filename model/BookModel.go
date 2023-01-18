package model

type Book struct {
	UUID        int    `json:"uuid"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishDate string `json:"publishDate"`
	ISBN        string `json:"ISBN"`
}
