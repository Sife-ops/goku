package utilities

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

var client *redis.Client
func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}
}
