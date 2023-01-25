package main

import (
	"fmt"
	"github.com/anisurahman75/apiDesign/api"
	"net/http"
)

func main() {
	// create router by calling CreateRouter function
	s := api.CreateNewServer()
	s.MountHandlers()
	err := http.ListenAndServe(":3000", s.Router)
	if err != nil {
		fmt.Printf("error : %s\n", err.Error())
	}
}
