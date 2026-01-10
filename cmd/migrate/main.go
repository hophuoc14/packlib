package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"packlib/config"
	"packlib/db"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	// Load configuration
	configuration := config.New()
	dbHost := configuration.Get("DATABASE_HOST")
	dbPort := configuration.Get("DATABASE_PORT")
	dbUser := configuration.Get("DATABASE_USERNAME")
	dbPassword := configuration.Get("DATABASE_PASSWORD")
	dbName := configuration.Get("DATABASE_NAME")

	// Create database connection using database/sql (not GORM)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer sqlDB.Close()

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Execute command
	switch command {
	case "up":
		if err := db.RunMigrations(sqlDB, dbName); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("✓ Migrations applied successfully")

	case "down":
		if err := db.RollbackMigration(sqlDB, dbName); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		log.Println("✓ Rolled back successfully")

	case "version":
		version, dirty, err := db.GetVersion(sqlDB, dbName)
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		fmt.Printf("Current version: %d\n", version)
		fmt.Printf("Dirty: %t\n", dirty)

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Please provide version number: go run cmd/migrate/main.go force <version>")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		if err := db.ForceVersion(sqlDB, dbName, version); err != nil {
			log.Fatalf("Force version failed: %v", err)
		}
		log.Printf("✓ Forced version to %d", version)

	default:
		log.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Database Migration Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  go run cmd/migrate/main.go <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  up              Apply all pending migrations")
	fmt.Println("  down            Rollback the last migration")
	fmt.Println("  version         Show current migration version")
	fmt.Println("  force <version> Force set migration version (use with caution)")
	fmt.Println("\nExamples:")
	fmt.Println("  go run cmd/migrate/main.go up")
	fmt.Println("  go run cmd/migrate/main.go down")
	fmt.Println("  go run cmd/migrate/main.go version")
	fmt.Println("  go run cmd/migrate/main.go force 1")
}
