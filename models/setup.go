package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	_, err = db.DB()

	if err != nil {
		log.Fatalf("DB Connection Failed : %v", err)
	}
	db.AutoMigrate(&Student{})

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	// Ping
	sqlDB.Ping()

	// Close
	sqlDB.Close()

	// Returns database statistics
	sqlDB.Stats()

}
