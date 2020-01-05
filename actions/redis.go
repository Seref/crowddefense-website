package actions

import (
	"log"

	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "wailord:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Printf("Errror initializing Redis Client %v", err)
	}
	log.Printf("%v", pong)
}