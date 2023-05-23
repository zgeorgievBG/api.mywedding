package database

import (
	"log"

	"api.mywedding/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error
	dsn := "host=127.0.0.1 user=postgres password=root dbname=mywedding port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to DB")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Artist{})
	// DB.Exec("ALTER TABLE artists ADD CONSTRAINT fk_artists_users FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE")
	log.Println("Migrated DB")
}
