package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	// "strconv"
	_ "github.com/lib/pq"
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
    // DB.Create(&Product{Name: "Steak", Price: 500})

	router.GET("/products", getProducts)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)
	router.Run("0.0.0.0:80")
}

func getProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	fmt.Println(products)
	c.IndentedJSON(http.StatusOK, products)
}