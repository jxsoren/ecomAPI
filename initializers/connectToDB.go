package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// DB Environment Variables
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// DB Connection String
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	
	// Connect to DB using DSN
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	// Try to get underlying MySQL db
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	// Try to ping DB to verify connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	fmt.Println("Connected to DB")
}
