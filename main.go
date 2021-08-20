package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	// "strconv"
	// "github.com/lib/pq"
	"whatsinme-api/models"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type Product struct {
  gorm.Model
  ID	uint	`json:"id" gorm:"primary_key"`
  Name  string	`json:"name"`
  Price uint 	`json:"price"`
}

func main() {

	router := gin.Default()

	models.OpenConnection();
    models.DB.Create(&models.Login{Email: "bill@gates.com", Password: "pass"})
    models.DB.Create(&models.Login{Email: "bill@gates.com", Password: "pass"})
    models.DB.Create(&models.Login{Email: "bill@gates.com", Password: "pass"})
    models.DB.Create(&models.Login{Email: "bill@gates.com", Password: "pass"})
    models.DB.Create(&models.Login{Email: "bill@gates.com", Password: "pass"})

	router.GET("/products", getProducts)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)
	router.Run("0.0.0.0:80")
}

func getProducts(c *gin.Context) {
	var products []Product
	models.DB.Find(&products)
	fmt.Println(products)
	c.IndentedJSON(http.StatusOK, products)
}