# Van Rakshak [![Build Status](https://api.travis-ci.org/thekosmix/gir-rakshak.png)](http://travis-ci.org/gir-rakshak)
Backend in Go for Van Rakshak App

## How to setup
### Prerequisite
1. mysql - for storing related data
2. go - for running backend server
3. redis - for storing auth tokens
4. postman - for testing the APIs
4. Any editor (VS code is prefferred)
### Setup and run
1. After cloning this repo, change mysql credentails in repo.go and redis credentials in redis.go
2. Create db gir_rakshak in mysql
3. Import gir_rakshak.sql file in gir_rakshak schema
4. Import [postman collection](https://www.getpostman.com/collections/5aba816fefaa2d8d2fd4)
5. Run command `go run main.go`

### Login and other activities
1. Open postman collections and start with ``User - Login``, which will give an access_token (at) and user_id (uid)
2. For all other authorised APIs, send this combo in headers
