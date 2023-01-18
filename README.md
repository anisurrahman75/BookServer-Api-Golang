# BookServer-Api-Golang
# How JWT Authorization works
    1. Fist User logged-in by userName & Password
    2. then server generate a jwt token( header,payload && Signatue)
    3. Header= header contains Encryption algorithm & type
    4. Payload= contains user defines information
    5. Signature= Server generate signature with server secret key
    6. when same user trying to do visit any end-point where need JWT Authorization..
    7. then we need to pass authorization request with bearer token
    8. Server will check this bearer token and verify it.
# Resources
    1.JWT:https://blog.logrocket.com/jwt-authentication-go/ 
    