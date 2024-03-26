package models

// User represents a user in the system
// @Description User object contains user information.
type User struct {
	// We hide the id because we don't want it to leave beyond the context of the database
	Id *string `json:"-"` // This field will be ignored by Swagger as it's unexported (private)
	// HashId is the public identifier for the user
	HashId *string `json:"id"` // HashId is a UUID
	// UserName is the user's username, required, between 5 and 50 characters
	UserName *string `json:"user_name" validate:"required,min=5,max=50"`
	// FirstName is the user's first name, required, between 5 and 50 characters
	FirstName *string `json:"first_name" validate:"required,min=3,max=255"`
	// LastName is the user's last name, required, between 5 and 50 characters
	LastName *string `json:"last_name" validate:"required,min=3,max=255"`
	// Email is the user's email address, required, must be a valid email format
	Email *string `json:"email" validate:"required,email,max=255"`
	// UserStatus is the user's status, required, must be exactly 1 character and contain "IAT"
	UserStatus *string `json:"user_status" validate:"required,oneof=I A T"`
	// Department is the user's department, can be null
	Department *string `json:"department"`
}
