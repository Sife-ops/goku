package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"whatsinme-api/models"
	"whatsinme-api/controllers"
)

func main() {
	router := gin.Default()

	models.OpenConnection()
	// models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})

	router.GET("/accounts", controllers.GetAccounts)
	router.GET("/accounts/:id", controllers.GetAccount)
	router.POST("/accounts", controllers.PostAccount)

	router.Run("0.0.0.0:" + os.Getenv("APP_PORT"))
}