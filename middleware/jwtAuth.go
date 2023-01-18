package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var sampleSecretKey = []byte("bolaJabeNah")

func GenerateJWT(userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = userName
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authHeader := request.Header.Get("Authorization")
		if len(authHeader) > 0 {
			jwtToken := authHeader[7:]
			token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected sighning method: %v", t.Header["alg"])
				}
				return sampleSecretKey, nil
			})

			// parsing errors result
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := writer.Write([]byte(err.Error()))
				if err2 != nil {
					return
				}
			}
			// if there's a token
			if token.Valid {
				next.ServeHTTP(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to Invalid Token"))
				if err != nil {
					return
				}
			}
		} else {
			// response for if there's no token header
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header\n Please log-in and get token first"))
			if err != nil {
				return
			}
		}
	})
}
