package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"))
}

// InitDB initializes the GORM DB connection.
func InitDB() {
	var err error
	connectStr := GetPostgresConnectionString()
	fmt.Println("Connecting to:", connectStr)

	db, err = gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Automatically migrate your schema (if your model struct has the necessary tags)
	// db.AutoMigrate(&model.UserDTO{})
}

// GetUser retrieves a user by its HASH column.
func GetUser(hashId string) (*model.UserDTO, error) {
	var user model.UserDTO
	// Ensure the table name is correctly referenced (if needed, use Table("USERS"))
	if err := db.Table("users").Where("HASH = ?", hashId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers retrieves a page of users that match a search criteria.
func GetUsers(search string, page, limit int) (*model.UserDTOPaging, error) {
	var users []model.UserDTO
	var totalCount int64

	query := db.Table("users")
	if search != "" {
		likeSearch := "%" + search + "%"
		// Using ILIKE for case-insensitive matching in PostgreSQL.
		query = query.Where("USER_NAME ILIKE ? OR FIRST_NAME ILIKE ? OR LAST_NAME ILIKE ? OR EMAIL ILIKE ?", likeSearch, likeSearch, likeSearch, search+"%")
	}

	// Count total matching records.
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	// Retrieve the users with pagination.
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	// Convert totalCount to int (if your model expects an int pointer).
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
	id := uuid.New().String()
	requestedUser.Id = &id

	// Compute a hash for the user.
	hash, err := model.HashObject(requestedUser)
	if err != nil {
		return nil, err
	}
	requestedUser.HashId = hash

	// Insert the record into the USERS table.
	if err := db.Table("USERS").Create(&requestedUser).Error; err != nil {
		return nil, err
	}
	return &requestedUser, nil
}

// UpdateUser updates an existing user's details.
func UpdateUser(requestedUser model.UserDTO) (*model.UserDTO, error) {
	var user model.UserDTO
	// Retrieve the user to be updated by its HASH.
	if err := db.Table("USERS").Where("HASH = ?", *requestedUser.HashId).First(&user).Error; err != nil {
		return nil, err
	}
	if user.Id == nil || *user.Id == "" {
		return nil, errors.New("user not found")
	}

	// Update fields if provided.
	if requestedUser.Department != nil && *requestedUser.Department != "" {
		user.Department = requestedUser.Department
	}
	if requestedUser.FirstName != nil && *requestedUser.FirstName != "" {
		user.FirstName = requestedUser.FirstName
	}
	if requestedUser.LastName != nil && *requestedUser.LastName != "" {
		user.LastName = requestedUser.LastName
	}
	if requestedUser.Email != nil && *requestedUser.Email != "" {
		user.Email = requestedUser.Email
	}
	if requestedUser.UserStatus != nil && *requestedUser.UserStatus != "" {
		user.UserStatus = requestedUser.UserStatus
	}

	// Recompute the hash after updates.
	hash, err := model.HashObject(user)
	if err != nil {
		return nil, err
	}
	user.HashId = hash

	// Update the record in the USERS table.
	if err := db.Table("USERS").Where("ID = ?", *user.Id).Updates(user).Error; err != nil {
		return nil, err
	}

	// Optionally, re-fetch the updated record.
	if err := db.Table("USERS").Where("ID = ?", *user.Id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser removes a user record based on its HASH.
func DeleteUser(id string) error {
	if err := db.Table("USERS").Where("HASH = ?", id).Delete(&model.UserDTO{}).Error; err != nil {
		return err
	}
	return nil
}
