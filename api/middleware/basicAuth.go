package middleware

import (
	"fmt"
	"net/http"
)

func checkValidUser(userName, password string) bool {
	return false
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {
			if !checkValidUser(username, password) {
				http.Error(w, "Wrong UserName or Password\n", http.StatusUnauthorized)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			_, err := w.Write([]byte("Login Successfully\n"))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
	})
}
