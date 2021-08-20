package main

import (
	// "strconv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

var DB *gorm.DB
func OpenConnection() {
	dsn := "host=207.246.94.25 user=postgres password=postgres dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed" + db.Name());
	}
	// db.AutoMigrate(&Product{})
	DB = db
}

func main() {

	router := gin.Default()

	OpenConnection();
    DB.Create(&Product{Name: "Steak", Price: 500})
    DB.Create(&Product{Name: "Steak", Price: 500})
    DB.Create(&Product{Name: "Steak", Price: 500})
    DB.Create(&Product{Name: "Steak", Price: 500})
    DB.Create(&Product{Name: "Steak", Price: 500})
    DB.Create(&Product{Name: "Steak", Price: 500})

	router.GET("/products", getProducts)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("0.0.0.0:80")
}

func getProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	fmt.Println(products)
	c.IndentedJSON(http.StatusOK, products)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
