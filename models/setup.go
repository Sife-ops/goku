package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func OpenConnection() {
	dsn := "host=207.246.94.25 user=postgres password=postgres dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed" + db.Name());
	}
	db.AutoMigrate(&Login{})
	DB = db
	// return db
}
