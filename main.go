package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"whatsinme-api/models"
)

func main() {

	router := gin.Default()

	models.OpenConnection()
	// models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})

	router.GET("/accounts", getAccounts)
	router.GET("/accounts/:id", getAccount)
	router.POST("/accounts", postAccount)

	router.Run("0.0.0.0:80")
}

func getAccounts(c *gin.Context) {
	var accounts []models.Account
	models.DB.Find(&accounts)
	fmt.Println(accounts)
	c.IndentedJSON(http.StatusOK, accounts)
}

func getAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	models.DB.First(&account, id)
	c.IndentedJSON(http.StatusOK, account)
}

func postAccount(c *gin.Context) {
	var newAccount models.Account
	if err := c.BindJSON(&newAccount); err != nil {
		return
	}
	models.DB.Create(&newAccount)
}