package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB(url string) {
	var err error
	if db, err = gorm.Open(sqlite.Open(url), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connected")
	// preparing db
	db.AutoMigrate(&customer{})
	log.Println("DB ready")
}
