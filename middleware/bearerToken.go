package middleware

func BearerToken(userName string) string {
	str, _ := GenerateJWT(userName)
	return "Bearer " + str
}
