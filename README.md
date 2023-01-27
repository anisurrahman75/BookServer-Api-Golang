# BookServer-Api-Golang
This API server provides endpoints to create,read,update & delete users and Books.
***
# Data Model Json
***
- User Model:
```
{
    "firstName": "Mridul",
    "lastName": "Halder",
    "userName": "mridul12",
    "password": "123"
}
```
- Book Model:
```
{
    "uuid": 11,
    "name": "learn-api",
    "author": "Anisur",
    "publishDate": "01-02-2022",
    "ISBN": "0999-0555-5954"
}
```
# To Start API Server Locally
***
`$ git clone https://github.com/anisurrahman75/BookServer-Api-Golang.git`

`$ git mod tidy && go mod vendor`

`$ go run main.go`
# To Start API Server Locally
***
`$ git clone https://github.com/anisurrahman75/BookServer-Api-Golang.git`

`$ git mod tidy && go mod vendor`

`$ go build`

`$ ./apiDesign startServer` 

**Note: default port:3030 && Authentication: JWT**
## Custom Port && Without Authentication ##
`$ ./apiDesign startServer -p=<your_port> -a=false`
# Available API EndPoints

Method | API EndPoint        | Authentication Type   | PayLoad               | Description                                             |Curl Command                                      
--- |---------------------|-----------------------|-----------------------|---------------------------------------------------------|-------------------
 GET| /api                | No-Auth               | Not-Required          | Welcome page of this Book-Server                        |`$ curl -X GET http://localhost:3000/api`
POST| /api/registerUser   | No-Auth               | Json File: User Model | Register a new User                                     |`$ curl -X POST -H "Content-Type:application/json" -d '<userModelJson>' http://localhost:3000/api/registerUser`
POST| /api/logIn          | Basic-Auth            | Not-Required          | LogIn with registered user,pass and get bearer token    |`$ curl -X POST --user  '<userName>:<passWord>' localhost:3000/api/logIn`
 GET| /api/books          | Bearer Token Required | Not-Required          | Get all BooksList frrom database in response            |`$ curl -X GET -H "Authorization: Bearer <bearerToken>" http://localhost:3000/api/books`
POST| /api/books          | Bearer Token Required | Json File: Book Model | Add new book and return added bookDetails in response   |`$ curl -X POST -H "Authorization: Bearer <bearerToken>" -H "Content-Type:application/json" -d '<bookModelJson>' localhost:3000/api/books`
 GET| /api/books/{bookId} | Bearer Token Required | Not-Required          | Search book with bookId and get bookDetails in response |`$ curl -X GET -H "Authorization: Bearer <bearerToken>" localhost:3000/api/books/<bookId>`
 PUT| /api/books/{bookId} | Bearer Token Required | Json File: Book Model | Update book and return updated bookDetails in response  |`$ curl -X PUT -H "Authorization: Bearer <bearerToken>" -H "Content-Type:application/json" -d '<bookModelJson>' localhost:3000/api/books/<bookId>`
DELETE| /api/books/{bookId} | Bearer Token Required | Not-Required          | Delete a aook and return deleted bookDetails in response|`$ curl -X DELETE -H "Authorization: Bearer <bearerToken>" localhost:3000/api/books/<bookId>`

# Resources
***
* [sysdevbd learn GO](https://sysdevbd.com/go/)
* [A Beginner’s Guide to HTTP and REST](https://code.tutsplus.com/tutorials/a-beginners-guide-to-http-and-rest--net-16340)
* [HTTP Response Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
* [JWT Token](https://blog.logrocket.com/jwt-authentication-go)
* [UNIT Testing](https://go-chi.io/#/pages/testing)