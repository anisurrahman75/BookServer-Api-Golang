# BookServer-Api-Golang
# Data Model
***
- User Model:
```
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
```
- Book Model:
```
type Book struct {
	UUID        int    `json:"uuid"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishDate string `json:"publishDate"`
	ISBN        string `json:"ISBN"`
}
```
# Available API EndPoints
***
Method | API EndPoint        | Authentication Type   | PayLoad               | Description                                              |Curl Command                                      
--- |---------------------|-----------------------|-----------------------|----------------------------------------------------------|-------------------
 GET| /api                | No-Auth               | Not-Required          | Welcome page of this Book-Server                         |``````curl -X GET localhost:3000/api``````
POST| /api/registerUser   | No-Auth               | Json File: User Model | Register a new User                                      |
POST| /api/logIn          | Basic-Auth            | Not-Required          | LogIn with registered user,pass and get bearer token     |
 GET| /api/books          | Bearer Token Required | Not-Required          | Get all BooksList frrom database in response             |
POST| /api/books          | Bearer Token Required | Json File: Book Model | Add new book and return added bookDetails in response    |
 GET| /api/books/{bookId} | Bearer Token Required | Not-Required          | Search book with bookId and get bookDetails in response  |
 PUT| /api/books/{bookId} | Bearer Token Required | Json File: Book Model | Update book and return updated bookDetails in response   |
DELETE| /api/books/{bookId} | Bearer Token Required | Not-Required          | Delete a aook and return deleted bookDetails in response |
# Sample Curl  Command:
***
> **Start Server**: download this repo. Inside repo directory run cmd: ```go run main.go```

> **Welcome Page**: ```curl -X GET localhost:3000/api```




# How JWT Authorization works
    1. user register himself by given by registration process
    2. At fist User logged-in by userName & Password
    2. then server generate a jwt token( header,payload && Signatue)
    3. Header= header contains Encryption algorithm & type
    4. Payload= contains user defines information
    5. Signature= Server generate signature with server secret key
    6. when any  user trying to do visit any end-point where need JWT Authorization..
    7. then user need to pass authorization request with bearer token
    8. Server will check this bearer token and verify it.
# Resources
    1.JWT:https://blog.logrocket.com/jwt-authentication-go/ 
    