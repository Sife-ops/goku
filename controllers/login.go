package controllers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"whatsinme-api/models"
)

// var hmacSecret []byte
var hmacSecret = []byte("temp")

func Login(c *gin.Context) {
	var credentials models.Account
	if err := c.BindJSON(&credentials); err != nil {
		return
	}

	var account models.Account
	models.DB.Where("email = ?", credentials.Email).First(&account)
	if credentials.Password == account.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, err := token.SignedString(hmacSecret)
		if err != nil {
			return
		}
		// fmt.Printf("_%s_\n", tokenString)
		c.IndentedJSON(http.StatusOK, tokenString)
	} else {
		return
		// c.IndentedJSON(http.StatusUnauthorized, "false")
	}
}
