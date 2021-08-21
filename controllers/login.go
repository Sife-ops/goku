package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "whatsinme-api/models"
	. "whatsinme-api/utilities"
)

func Login(c *gin.Context) {
	var credentials Account
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided.")
		return
	}

	var account Account
	DB.Where("email = ?", credentials.Email).First(&account)

	if credentials.Email != account.Email ||
		credentials.Password != account.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login credentials.")
		return
	}

	token, err := CreateToken(account.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
}
