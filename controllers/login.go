package controllers

import (
	"net/http"
	"os"
	// "time"
	"whatsinme-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {
	var credentials models.Account
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided.")
		return
	}

	var account models.Account
	models.DB.Where("email = ?", credentials.Email).First(&account)

	//compare the user from the request, with the one we defined:
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

func CreateToken(accountId uint) (string, error) {
	var err error
	
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["account_id"] = accountId
	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
