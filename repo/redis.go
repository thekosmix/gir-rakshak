package repo

import (
	"log"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v3"
)

var cache *bigcache.BigCache

var tokenExpiryTime = 24 * time.Hour
var tokenPrefix = "gir:user:"

func InitCache() {
	// Initialize the bigcache instance
	config := bigcache.DefaultConfig(tokenExpiryTime)
	config.CleanWindow = 1 * time.Hour // Adjust the cleanup interval as needed

	var err error
	cache, err = bigcache.NewBigCache(config)
	if err != nil {
		log.Printf("Error initializing cache: %v", err)
	}
}

func AddUserToken(id int, sessionToken string) error {
	key := tokenPrefix + strconv.Itoa(id)
	return cache.Set(key, []byte(sessionToken))
}

func GetUserToken(id int) string {
	key := tokenPrefix + strconv.Itoa(id)
	token, err := cache.Get(key)
	if err != nil {
		log.Printf("Error retrieving token: %v", err)
		return ""
	}
	return string(token)
}
