package controllers

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"whatsinme-api/models"
)

var hmacSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

func Login(c *gin.Context) {
	var credentials models.Account
	if err := c.BindJSON(&credentials); err != nil {
		return
	}

	var account models.Account
	models.DB.Where("email = ?", credentials.Email).First(&account)
	if credentials.Password == account.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": account.Email,
		})
		tokenString, err := token.SignedString(hmacSecret)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusOK, tokenString)
	} else {
		// send back WWW-Authenticate header?
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
