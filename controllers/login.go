package controllers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"whatsinme-api/models"
)

func Login(c *gin.Context) {
	var credentials models.Account
	if err := c.BindJSON(&credentials); err != nil {
		return
	}

	var account models.Account
	models.DB.Where("email = ?", credentials.Email).First(&account)
	if credentials.Password == account.Password {
		c.IndentedJSON(http.StatusOK, "true")
	} else {
		c.IndentedJSON(http.StatusUnauthorized, "false")
	}
}
