package middleware

import (
	"fmt"
)

func BearerToken(userName string) string {
	str, err := GenerateJWT(userName)
	if err != nil {
		fmt.Println("Error on GenerateJWT token FROM barerToken Function\n")
		return err.Error()
	}
	return "Bearer " + str
}
