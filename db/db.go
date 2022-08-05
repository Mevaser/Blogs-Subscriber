package db

import (
	"blogs_subscriber/config"
	"blogs_subscriber/db/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	log.Printf("[*] Initializing %s.", config.Config.SqliteDb)
	db, err := gorm.Open(sqlite.Open(config.Config.SqliteDb), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Blog{})
}