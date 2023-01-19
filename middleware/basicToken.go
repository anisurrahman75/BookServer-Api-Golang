package middleware

import (
	"encoding/base64"
)

func BasicToken(user, pas string) string {
	str := user + ":" + pas
	token := "Basic " + base64.StdEncoding.EncodeToString([]byte(str))
	return token
}
