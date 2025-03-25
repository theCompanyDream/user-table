package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	model "github.com/theCompanyDream/user-angular/apps/backend/models"
)

var db *gorm.DB

// GetPostgresConnectionString returns a PostgreSQL connection string.
// It first checks for POSTGRES_URL and falls back to constructing the string.
func GetPostgresConnectionString() string {
	connectStr := os.Getenv("POSTGRES_URL")
	if connectStr != "" {
		return connectStr
	}

	// More explicit connection string for Docker/internal network
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"))
}

func InitDB() {
	var err error
	connectStr := GetPostgresConnectionString()
	fmt.Println("Connecting to:", connectStr)

	// Add more verbose logging and configuration
	db, err = gorm.Open(postgres.Open(connectStr), &gorm.Config{
		// Add additional configurations
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logging
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database: %v", err)
	}

	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Auto migrate with more detailed error handling
	if err := db.AutoMigrate(&model.UserDTO{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	fmt.Println("Database connection successful")
}

// GetUser retrieves a user by its HASH column.
func GetUser(hashId string) (*model.UserDTO, error) {
	var user model.UserDTO
	// Ensure the table name is correctly referenced (if needed, use Table("users"))
	if err := db.Table("users").Where("HASH = ?", hashId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers retrieves a page of users that match a search criteria.
func GetUsers(search string, page, limit int, c echo.Context) (*model.UserDTOPaging, error) {
	var users []model.UserDTO
	var totalCount int64

	// Use db.Model instead of db.Table
	query := db.Model(&model.UserDTO{})

	if search != "" {
		likeSearch := "%" + search + "%"
		query = query.Where(
			"user_name ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?",
			likeSearch, likeSearch, likeSearch, likeSearch,
		)
	}

	// Count total matching records
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	// Remove explicit Select, let GORM handle field mapping
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	c.Logger().Info("Total users: %+v\n", users)

	total := int(totalCount)
	paging := model.Paging{
		Page:     &page,
		Length:   &total,
		PageSize: &limit,
	}
	return &model.UserDTOPaging{
		Paging: paging,
		Users:  users,
	}, nil
}

// CreateUser creates a new user record.
func CreateUser(requestedUser model.UserDTO) (*model.UserDTO, error) {
	// Generate a new UUID for the user.
	id := uuid.New()
	requestedUser.ID = id

	// Compute a hash for the user.
	hash, err := model.HashObject(requestedUser)
	if err != nil {
		return nil, err
	}
	requestedUser.Hash = *hash

	// Insert the record into the USERS table.
	if err := db.Table("users").Create(&requestedUser).Error; err != nil {
		return nil, err
	}
	return &requestedUser, nil
}

// UpdateUser updates an existing user's details.
func UpdateUser(requestedUser model.UserDTO) (*model.UserDTO, error) {
	var user model.UserDTO
	// Retrieve the user to be updated by its HASH.
	if err := db.Table("users").Where("hash LIKE ?", requestedUser.Hash).First(&user).Error; err != nil {
		return nil, err
	}
	if user.ID == uuid.Nil {
		return nil, errors.New("user not found")
	}

	// Update fields if provided.
	if requestedUser.Department != nil && *requestedUser.Department != "" {
		user.Department = requestedUser.Department
	}
	if requestedUser.FirstName != "" {
		user.FirstName = requestedUser.FirstName
	}
	if requestedUser.LastName != "" {
		user.LastName = requestedUser.LastName
	}
	if requestedUser.Email != "" {
		user.Email = requestedUser.Email
	}
	if requestedUser.UserStatus != "" {
		user.UserStatus = requestedUser.UserStatus
	}

	// Recompute the hash after updates.
	hash, err := model.HashObject(user)
	if err != nil {
		return nil, err
	}
	user.Hash = *hash

	// Update the record in the USERS table.
	if err := db.Table("users").Where("ID = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}

	// Optionally, re-fetch the updated record.
	if err := db.Table("users").Where("ID = ?", user.ID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser removes a user record based on its HASH.
func DeleteUser(id string) error {
	if err := db.Table("users").Where("HASH = ?", id).Delete(&model.UserDTO{}).Error; err != nil {
		return err
	}
	return nil
}
