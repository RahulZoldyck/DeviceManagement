package main

import (
	"device-management-api/config"
	"device-management-api/internal/models"
	"device-management-api/internal/routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set up MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Connect to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the Device model (creates the table if it doesn’t exist)
	db.AutoMigrate(&models.Device{})

	// Set up router and pass the database connection
	r := routes.SetupRouter(db)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
