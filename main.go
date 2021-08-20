package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"net/http"
	// "strconv"
	// "github.com/lib/pq"
	"whatsinme-api/models"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// type Product struct {
// 	gorm.Model
// 	ID    uint   `json:"id" gorm:"primary_key"`
// 	Name  string `json:"name"`
// 	Price uint   `json:"price"`
// }

func main() {

	router := gin.Default()

	models.OpenConnection()
	models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})
	models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})
	models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})
	models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})
	models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})

	router.GET("/accounts", getAccounts)
	router.GET("/accounts/:id", getAccount)
	// router.POST("/accounts", createAccount)

	// router.GET("/products", getProducts)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)
	
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

	// // Loop over the list of albums, looking for
	// // an album whose ID value matches the parameter.
	// for _, a := range albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// func getProducts(c *gin.Context) {
// 	var products []Product
// 	models.DB.Find(&products)
// 	fmt.Println(products)
// 	c.IndentedJSON(http.StatusOK, products)
// }
