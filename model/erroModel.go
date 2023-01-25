package model

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	status     string
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	ErrorMsg   string `json:"errorMsg"`
}

func RequestForError(statusCode int, err string, w http.ResponseWriter, msg string) {
	temp := Error{
		"----------FAILED------------",
		statusCode,
		err,
		msg,
	}
	w.WriteHeader(statusCode)
	er := json.NewEncoder(w).Encode(temp)
	http.Error(w, er.Error(), http.StatusBadRequest)
}
