package utilities

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	//Initializing redis
	// dsn := os.Getenv("REDIS_DSN")
	// if len(dsn) == 0 {
	dsn := "localhost:6379"
	// }
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}
}
