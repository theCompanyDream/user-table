package models

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
	HashId *string `json:"id" param:"id"` // HashId is a UUID
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
	Id *string `json:"-"` // This field will be ignored by Swagger as it's unexported (private)
	// HashId is the public identifier for the user
	HashId *string `json:"id"` // HashId is a UUID
	// UserName is the user's username, between 5 and 50 characters
	UserName *string `json:"user_name"`
	// FirstName is the user's first name, between 5 and 50 characters
	FirstName *string `json:"first_name"`
	// LastName is the user's last name, between 5 and 50 characters
	LastName *string `json:"last_name"`
	// Email is the user's email address, must be a valid email format
	Email *string `json:"email"`
	// UserStatus is the user's status, must be exactly 1 character and contain "IAT"
	UserStatus *string `json:"user_status"`
	// Department is the user's department, can be null
	Department *string `json:"department"`
}

func CreateToDTO(userCreate UserCreate) UserDTO {
	return UserDTO{
		Id:         userCreate.Id,
		HashId:     userCreate.HashId,
		UserName:   userCreate.UserName,
		FirstName:  userCreate.FirstName,
		LastName:   userCreate.LastName,
		Email:      userCreate.Email,
		UserStatus: userCreate.UserStatus,
		Department: userCreate.Department,
	}
}

func UpdateToDTO(userUpdate UserUpdate) UserDTO {
	return UserDTO{
		Id:         userUpdate.Id,
		HashId:     userUpdate.HashId,
		UserName:   userUpdate.UserName,
		FirstName:  userUpdate.FirstName,
		LastName:   userUpdate.LastName,
		Email:      userUpdate.Email,
		UserStatus: userUpdate.UserStatus,
		Department: userUpdate.Department,
	}
}
