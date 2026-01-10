package config

import (
	"fmt"
	// "log"
	// dbmigrate "packlib/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func NewPostgresDatabase(configuration Config) (*gorm.DB, error) {
	dbHost := configuration.Get("DATABASE_HOST")
	dbPort := configuration.Get("DATABASE_PORT")
	dbUser := configuration.Get("DATABASE_USERNAME")
	dbPassword := configuration.Get("DATABASE_PASSWORD")
	dbName := configuration.Get("DATABASE_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	database = db

	// Run migrations
	_, err = db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Uncomment the following lines to run migrations automatically on startup
	// You can also run migrations manually using: go run cmd/migrate/main.go up
	/*
	if err := dbmigrate.RunMigrations(sqlDB, dbName); err != nil {
		log.Printf("Warning: Migration failed: %v", err)
		// Decide whether to fail startup or continue
		// return nil, err
	}
	*/

	return database, nil
}

func GetDatabase() *gorm.DB {
	return database
}

