package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"whatsinme-api/models"
)

func GetAccounts(c *gin.Context) {
	var accounts []models.Account
	models.DB.Find(&accounts)
	fmt.Println(accounts)
	c.IndentedJSON(http.StatusOK, accounts)
}

func GetAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	models.DB.First(&account, id)
	c.IndentedJSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	id := c.Param("id")

	var account models.Account
	models.DB.First(&account, id)

	var newAccount models.Account
	if err := c.BindJSON(&newAccount); err != nil {
		return
	}

	if newAccount.Email != "" ||
		newAccount.Password != "" {
		if newAccount.Email != "" {
			account.Email = newAccount.Email
		}
		if newAccount.Password != "" {
			account.Password = newAccount.Password
		}
		models.DB.Save(&account)
	}

	c.IndentedJSON(http.StatusOK, account)
}

func PostAccount(c *gin.Context) {
	var newAccount models.Account
	if err := c.BindJSON(&newAccount); err != nil {
		return
	}
	models.DB.Create(&newAccount)
}
