package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v \n", err)
	}

	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v \n", err)
	}

}
