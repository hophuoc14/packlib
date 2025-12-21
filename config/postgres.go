package config

import (
	"fmt"
	"packlib/entity"

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


	database.AutoMigrate(&entity.User{})
	database.AutoMigrate(&entity.Employee{})
	database.AutoMigrate(&entity.Department{})
	database.AutoMigrate(&entity.Project{})

	return database, nil
}

func GetDatabase() *gorm.DB {
	return database
}

