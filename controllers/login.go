package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "goku/models"
	. "goku/utilities"
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

	tokens, err := CreateToken(account.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := CreateAuth(account.ID, tokens)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"access_token": tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
