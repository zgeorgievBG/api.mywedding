package database

import (
	"log"

	"api.mywedding/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbErr error

func Connect(connectionString string) {
	DB, dbErr = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to DB")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Artist{})
	// DB.AutoMigrate(&models.Facilities{})
	log.Println("Migrated DB")
}
