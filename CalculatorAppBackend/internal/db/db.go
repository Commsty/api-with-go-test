package db

import (
	"calc/internal/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=starlyvtb3 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could connect to database: %v", err)
	}

	if err := db.AutoMigrate(&entity.Calculation{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return db, nil
}
