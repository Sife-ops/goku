package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"goku/controllers"
	"goku/models"
	"os"
)

func main() {
	router := gin.Default()

	models.OpenConnection()
	// models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})

	router.GET("/accounts", controllers.GetAccounts)
	router.GET("/accounts/:id", controllers.GetAccount)
	router.POST("/accounts", controllers.PostAccount)
	router.PUT("/accounts/:id", controllers.UpdateAccount)
	router.DELETE("/accounts/:id", controllers.DeleteAccount)

	router.POST("/login", controllers.Login)

	router.Run("0.0.0.0:" + os.Getenv("DOCKER_PORT"))
}

// EXPERIMENTAL
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
// EXPERIMENTAL
