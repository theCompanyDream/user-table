package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	model "github.com/theCompanyDream/user-angular/apps/backend/models"
)

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

func GetUser(requestedUser *model.User) (*model.User, error) {
	// Assume URL like /users/{id}
	var user *model.User
	query := squirrel.Select("HASHID, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT").
		From("users").
		Where(squirrel.Eq{"HASH": requestedUser.HashId}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	err := query.QueryRow().Scan(user.HashId, user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus, user.Department)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers(search string, page, limit int) ([]model.User, error) {
	// Assume URL like /users/{id}
	offset := (page - 1) * limit
	users := []model.User{}

	query := squirrel.Select("HASHID, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT").
		From("users").
		Where("USER_NAME LIKE ?", "%"+search+"%").
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
		var user model.User
		if err := rows.Scan(&user.HashId, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(requestedUser model.User) (*model.User, error) {
	// Parse user details from the request body and insert into the database
	var user *model.User
	id := uuid.New().String()
	requestedUser.Id = &id
	hash, err := model.HashObject(requestedUser)
	if err != nil {
		return nil, err
	}
	requestedUser.HashId = hash
	query := squirrel.Insert("users").
		Columns("ID", "HASH", "USER_NAME", "FIRST_NAME", "LAST_NAME", "EMAIL", "USER_STATUS", "DEPARTMENT").
		Values(requestedUser.Id, requestedUser.HashId, requestedUser.UserName, requestedUser.FirstName, requestedUser.LastName, requestedUser.Email, requestedUser.UserStatus, requestedUser.Department).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)
	_, err = query.Exec()
	if err != nil {
		return nil, err
	}
	err = query.QueryRow().Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(requestedUser model.User) (*model.User, error) {
	// Parse user details from the request body and insert into the database
	var user *model.User
	hash, err := model.HashObject(requestedUser)
	if err != nil {
		return nil, err
	}

	requestedUser.HashId = hash
	query := squirrel.Update("users").
		Set("HASH", requestedUser.HashId).
		Set("FIRST_NAME", requestedUser.FirstName).
		Set("LAST_NAME", requestedUser.LastName).
		Set("EMAIL", requestedUser.Email).
		Set("USER_STATUS", requestedUser.UserStatus).
		Set("DEPARTMENT", requestedUser.Department).
		Where(squirrel.Eq{"HASH": requestedUser.HashId}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	_, err = query.Exec()
	if err != nil {
		return nil, err
	}
	err = query.QueryRow().Scan(&user.Id, &user.HashId, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id string) error {
	// Parse user details from the request body and insert into the database
	query := squirrel.Delete("users").
		Where(squirrel.Eq{"HASH": id}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	_, err := query.Exec()
	if err != nil {
		return err
	}
	return nil
}
