package model

type User struct {
    ID    int    `json:"id"`
    FName  string `json:"first_name"`
	lName  string `json:"last_name"`
    Email string `json:"email"`
	user_status string `json:"user_status"`
	department string `json:"department"`
}
