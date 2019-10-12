package repo

import (
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn

var tokenExpiryTime = "86400"
var tokenPrefix = "gir:user:"

func InitCache() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		log.Printf(err.Error())
	}
	// Assign the connection to the package level `cache` variable
	cache = conn
}

func AddUserToken(id int, sessionToken string) (reply interface{}, err error) {
	return cache.Do("SETEX", tokenPrefix+strconv.Itoa(id), tokenExpiryTime, sessionToken)
}

func GetUserToken(id int) string {
	token, _ := redis.String(cache.Do("GET", tokenPrefix+strconv.Itoa(id)))
	return token
}
