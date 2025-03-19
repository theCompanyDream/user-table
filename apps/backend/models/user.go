package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

// User represents a user in the system
// @Description User object contains user information.
type UserCreate struct {
	// We hide the id because we don't want it to leave beyond the context of the database
	Id *string `json:"-"` // This field will be ignored by Swagger as it's unexported (private)
	// HashId is the public identifier for the user
	HashId *string `json:"id" param:"id"` // HashId is a UUID
	// UserName is the user's username, required, between 5 and 50 characters
	UserName *string `json:"user_name" validate:"required,min=5,max=50" form:"user_name"`
	// FirstName is the user's first name, required, between 5 and 50 characters
	FirstName *string `json:"first_name" validate:"required,min=3,max=255" form:"first_name"`
	// LastName is the user's last name, required, between 5 and 50 characters
	LastName *string `json:"last_name" validate:"required,min=3,max=255" form:"last_name"`
	// Email is the user's email address, required, must be a valid email format
	Email *string `json:"email" validate:"required,email,max=255" form:"email"`
	// UserStatus is the user's status, required, must be exactly 1 character and contain "IAT"
	UserStatus *string `json:"user_status" validate:"required,oneof=I A T" form:"user_status"`
	// Department is the user's department, can be null
	Department *string `json:"department" form:"department"`
}

type UserUpdate struct {
	Id *string `json:"-"` // This field will be ignored by Swagger as it's unexported (private)
	// HashId is the public identifier for the user
	HashId *string `json:"id" path:"id"` // HashId is a UUID
	// UserName is the user's username, between 5 and 50 characters
	UserName *string `json:"user_name" validate:"omitempty,min=5,max=50" form:"user_name"`
	// FirstName is the user's first name, between 5 and 50 characters
	FirstName *string `json:"first_name" validate:"omitempty,min=3,max=255" form:"first_name"`
	// LastName is the user's last name, between 5 and 50 characters
	LastName *string `json:"last_name" validate:"omitempty,min=3,max=255" form:"last_name"` // Corrected validate tag
	// Email is the user's email address, must be a valid email format
	Email *string `json:"email" validate:"omitempty,email,max=255" form:"email"`
	// UserStatus is the user's status, must be exactly 1 character and contain "IAT"
	UserStatus *string `json:"user_status" validate:"omitempty,oneof=I A T" form:"user_status"`
	// Department is the user's department, can be null
	Department *string `json:"department" form:"department"`
}

type UserDTO struct {
	// Id is the internal database identifier; omit from JSON output.
	Id uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"-"`
	// HashId is the public identifier for the user (varchar(64)).
	HashId string `gorm:"column:hash" json:"id"`
	// UserName is the user's username, between 5 and 50 characters.
	UserName string `gorm:"column:user_name" json:"user_name"`
	// FirstName is the user's first name, between 5 and 50 characters.
	FirstName string `gorm:"column:first_name" json:"first_name"`
	// LastName is the user's last name, between 5 and 50 characters.
	LastName string `gorm:"column:last_name" json:"last_name"`
	// Email is the user's email address, must be a valid email format.
	Email string `gorm:"column:email" json:"email"`
	// UserStatus is the user's status, must be exactly 1 character and contain "IAT".
	UserStatus string `gorm:"column:user_status" json:"user_status"`
	// Department is the user's department, can be null.
	Department *string `gorm:"column:department" json:"department"`
}

// TableName sets the table name for UserDTO to "users".
func (UserDTO) TableName() string {
	return "users"
}

func CreateToDTO(userCreate UserCreate) *UserDTO {
	var user UserDTO
	copier.Copy(&user, &userCreate)
	return &user
}

func UpdateToDTO(userUpdate UserUpdate) *UserDTO {
	var user UserDTO
	copier.Copy(&user, &userUpdate)
	return &user
}
