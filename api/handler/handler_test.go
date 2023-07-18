package handler

import (
	"bytes"
	middleware2 "github.com/anisurahman75/apiDesign/api/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Welcome(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	req, _ := http.NewRequest("GET", "/api", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}
func Test_logIn(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		expectedStatusCode int
		token              string
	}
	test := []Test{
		{"POST",
			"/api/logIn",
			nil,
			http.StatusOK,
			middleware2.BasicToken("sunny", "123")},
		{"POST",
			"/api/logIn",
			nil,
			http.StatusUnauthorized,
			middleware2.BasicToken("snny241",
				"123")},
		{"POST",
			"/api/logIn",
			nil,
			http.StatusOK,
			middleware2.BasicToken("mridul12", "123")},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		//fmt.Println(response.Code)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_Register(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		expectedStatusCode int
	}
	test := []Test{
		{
			"POST",
			"/api/registerUser",
			bytes.NewReader([]byte(`{"FirstName": "imtiaj", "LastName": "Halder", "UserName": "imtiaj12", "Password": "123"}`)),
			http.StatusOK,
		},
		{
			"POST",
			"/api/registerUser",
			bytes.NewReader([]byte(`{"FirstName": "imtiaj", "LastName": "Halder", "UserName": "imtiaj12, "Password": "123"}`)),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_AllBookList(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{
			"GET",
			"/api/books",
			nil,
			middleware2.BearerToken("testing"),
			200,
		},
		{
			"GET",
			"/api/books",
			nil,
			middleware2.BearerToken("testing"),
			200,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_AddBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"POST",
			"/api/books",
			bytes.NewReader([]byte(`{"UUID": 10, "Name": "learn-api", "Author": "Anisur", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5914"}`)),
			middleware2.BearerToken("sunny"),
			http.StatusOK,
		},
		{"POST",
			"/api/books",
			bytes.NewReader([]byte(`{"UUID: 10, "Name": "learn-api", "Author": "Anisur", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"}`)),
			middleware2.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		middleware2.VerifyJWT(s.Router)
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_FindBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"GET",
			"/api/books/1",
			nil,
			middleware2.BearerToken("sunny"),
			http.StatusOK,
		},
		{"GET",
			"/api/books/10",
			nil,
			middleware2.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_UpdateBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"PUT",
			"/api/books/4",
			bytes.NewReader([]byte(`{"UUID": 4, "Name": "learn-api", "Author": "RaselVai", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"},`)),
			middleware2.BearerToken("sunny"),
			http.StatusOK,
		},
		{"PUT",
			"/api/books/100",
			bytes.NewReader([]byte(`{"UUID": 4, "Name": "learn-api", "Author": "RaselVai", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"},`)),
			middleware2.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_DeleteBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"DELETE",
			"/api/books/4",
			nil,
			middleware2.BearerToken("sunny"),
			http.StatusOK,
		},
		{"DELETE",
			"/api/books/100",
			nil,
			middleware2.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
