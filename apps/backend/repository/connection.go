package repository

import (
	"fmt"
	"net/url"
	"os"

	model "github.com/theCompanyDream/user-table/apps/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// GetPostgresConnectionString returns a PostgreSQL connection string.
// It first checks for POSTGRES_URL and falls back to constructing the string.
func GetPostgresConnectionString() string {
	// More explicit connection string for Docker/internal network
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"))
}

func InitDB() error {
	var err error
	connectStr := GetPostgresConnectionString()
	fmt.Println("Connecting to:", connectStr)

	// Add more verbose logging and configuration
	db, err = gorm.Open(postgres.Open(connectStr), &gorm.Config{
		// Add additional configurations
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logging
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database: %v", err)
	}

	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Auto migrate with more detailed error handling
	if err := db.AutoMigrate(&model.UserDTO{}); err != nil {
		// Log the error as a warning and continue
		fmt.Printf("Warning: Failed to auto migrate: %v", err)
	}

	fmt.Println("Database connection successful")
	return nil
}

func modifyConnectionString(connStr string) (string, error) {
	parsed, err := url.Parse(connStr)
	if err != nil {
		return "", err
	}
	// Add or update the query parameter
	query := parsed.Query()
	query.Set("prefer_simple_protocol", "true")
	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}

func SeverlessInitDB() error {
	var err error
	connectStr := os.Getenv("POSTGRES_URL")
	if connectStr != "" {
		return fmt.Errorf("failed to retrieve connection string: POSTGRES_URL not set %v", err)
	}
	// Modify the connection string to disable prepared statement caching.
	modifiedConnStr, err := modifyConnectionString(connectStr)
	if err != nil {
		return fmt.Errorf("failed to modify connection string: %v", err)
	}
	fmt.Println("Connecting to:", modifiedConnStr)

	db, err = gorm.Open(postgres.Open(modifiedConnStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Database connection successful")
	return nil
}
