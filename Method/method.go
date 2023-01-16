package Method

import (
	"apiDesign/ServerDatabase"
	"encoding/json"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\nWelcome to My Book-Server"))
}
func AllBookList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ServerDatabase.BookList)
}
func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book ServerDatabase.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if _, ok := ServerDatabase.BookList[book.UUID]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("book with the id already exists"))
		return
	}
	w.Write([]byte("book successfully added. "))
	ServerDatabase.BookList[book.UUID] = book
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Println("failed to write data in response body")
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	return
}
