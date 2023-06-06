package database

import (
	"fmt"
	"log"

	"github.com/cglavin50/job-tracker/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the DB")
	}

	configureTable() // configure table to match our user model

	fmt.Println("DB connection established")
}

func configureTable() {
	err := DB.AutoMigrate(&models.Job{})
	if err != nil {
		log.Fatal("Error formatting table to model", err)
	}
}
