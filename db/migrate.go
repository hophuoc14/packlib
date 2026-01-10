package db

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// GetMigrationPath returns the absolute path to the migrations directory
func GetMigrationPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get current file path")
	}
	
	// Get the directory containing this file (db package)
	dir := filepath.Dir(filename)
	
	// migrations folder is in the same directory
	migrationsPath := filepath.Join(dir, "migrations")
	
	return migrationsPath
}

// NewMigrate creates a new migrate instance
func NewMigrate(db *sql.DB, databaseName string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	migrationsPath := GetMigrationPath()
	sourceURL := fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath))

	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		databaseName,
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return m, nil
}

// RunMigrations applies all pending migrations
func RunMigrations(db *sql.DB, databaseName string) error {
	m, err := NewMigrate(db, databaseName)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get migration version: %w", err)
	}

	if err == migrate.ErrNilVersion {
		log.Println("No migrations applied yet")
	} else {
		log.Printf("Current migration version: %d (dirty: %t)", version, dirty)
	}

	return nil
}

// RollbackMigration rolls back the last migration
func RollbackMigration(db *sql.DB, databaseName string) error {
	m, err := NewMigrate(db, databaseName)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	log.Println("Successfully rolled back last migration")
	return nil
}

// GetVersion returns the current migration version
func GetVersion(db *sql.DB, databaseName string) (uint, bool, error) {
	m, err := NewMigrate(db, databaseName)
	if err != nil {
		return 0, false, err
	}
	defer m.Close()

	return m.Version()
}

// ForceVersion sets the migration version without running migrations
// Use this with caution - only for recovery scenarios
func ForceVersion(db *sql.DB, databaseName string, version int) error {
	m, err := NewMigrate(db, databaseName)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Force(version); err != nil {
		return fmt.Errorf("failed to force version: %w", err)
	}

	log.Printf("Forced migration version to: %d", version)
	return nil
}
