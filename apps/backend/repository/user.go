package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type User struct {
	id         string
	HashId 	   string `json: "json:"Id"`
	UserName   string `json:"user_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UserStatus string `json:"user_status"`
	Department string `json:"department"`
}

var db *sql.DB

func GetPostgresConnectionString() string {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))

	return connectStr
}

func InitDB() {
	var err error
	connectStr := GetPostgresConnectionString()

	db, err = sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(requestedUser *User) (*User, error) {
	// Assume URL like /users/{id}
	var user *User
	query := squirrel.Select("*").
		From("users").
		Where(squirrel.Eq{"hash": requestedUser.id}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	err := query.QueryRow().Scan(&user.id, &user.Hash, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers(search string, page, limit int) ([]User, error) {
	// Assume URL like /users/{id}
	offset := (page - 1) * limit
	users := []User{}

	query := squirrel.Select("*").
		From("users").
		Where("first_name LIKE ?", "%"+search+"%").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.id, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(requestedUser *User) (*User, error) {
	// Parse user details from the request body and insert into the database
	var user *User
	query := squirrel.Insert("users").
		Columns("user_name", "first_name", "last_name", "email", "user_status", "department").
		Values(requestedUser.UserName, requestedUser.FirstName, requestedUser.LastName, requestedUser.Email, requestedUser.UserStatus, requestedUser.Department).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	_, err := query.Exec()
	if err != nil {
		return nil, err
	}
	err = query.QueryRow().Scan(&user.id, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(requestedUser User) (*User, error) {
	// Parse user details from the request body and insert into the database
	var user *User
	query := squirrel.Update("users").
		Set("first_name", requestedUser.FirstName).
		Set("last_name", requestedUser.LastName).
		Set("email", requestedUser.Email).
		Set("user_status", requestedUser.UserStatus).
		Set("department", requestedUser.Department).
		Where(squirrel.Eq{"id": requestedUser.id}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	_, err := query.Exec()
	if err != nil {
		return nil, err
	}
	err = query.QueryRow().Scan(&user.id, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) error {
	// Parse user details from the request body and insert into the database
	query := squirrel.Delete("users").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	_, err := query.Exec()
	if err != nil {
		return err
	}
	return nil
}
