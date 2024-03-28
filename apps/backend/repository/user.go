package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
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

func GetUser(hashId string) (*model.UserDTO, error) {
	// Assume URL like /users/{id}
	var user model.UserDTO
	query := squirrel.Select("Id, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT").
		From("USERS").
		Where(squirrel.Eq{"HASH": hashId}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	err := query.QueryRow().Scan(&user.Id, &user.HashId, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(search string, page, limit int) ([]model.UserDTO, error) {
	// Check for potential overflow during multiplication
	offset := uint64((page - 1) * limit)
	users := []model.UserDTO{}
	query := squirrel.Select("Id, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT").
		From("USERS")

	if search != "" {
		query = query.Where("USER_NAME LIKE ?", "%"+search+"%")
	}
	// Note: there was a weird bug that if offset was 0 it overflowed the buffer and made offset this obsurb number
	if offset != math.MaxUint64 {
		query = query.Offset(offset)
	}
	query = query.Limit(uint64(limit)).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UserDTO
		if err := rows.Scan(&user.Id, &user.HashId, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func CreateUser(requestedUser model.UserDTO) (*model.UserDTO, error) {
	var user model.UserDTO
	id := uuid.New().String()
	requestedUser.Id = &id
	hash, err := model.HashObject(requestedUser)
	if err != nil {
		return nil, err
	}
	requestedUser.HashId = hash

	query := squirrel.Insert("USERS").
		Columns("ID", "HASH", "USER_NAME", "FIRST_NAME", "LAST_NAME", "EMAIL", "USER_STATUS", "DEPARTMENT").
		Values(requestedUser.Id, requestedUser.HashId, requestedUser.UserName, requestedUser.FirstName, requestedUser.LastName, requestedUser.Email, requestedUser.UserStatus, requestedUser.Department).
		Suffix("RETURNING ID, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)

	err = query.QueryRow().Scan(&user.Id, &user.HashId, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(requestedUser model.UserDTO) (*model.UserDTO, error) {
	// Grab the user to be updated
	user, err := GetUser(*requestedUser.HashId)
	if err != nil {
		return nil, err
	} else if user.Id == nil && *user.Id == "" {
		return nil, errors.New("user not Found")
	}

	query := squirrel.Update("USERS")
	if requestedUser.Department != nil && *requestedUser.Department != "" {
		user.Department = requestedUser.Department
		query = query.Set("DEPARTMENT", *requestedUser.Department)
	}
	if requestedUser.FirstName != nil && *requestedUser.FirstName != "" {
		user.FirstName = requestedUser.FirstName
		query = query.Set("FIRST_NAME", *requestedUser.FirstName)
	}
	if requestedUser.LastName != nil && *requestedUser.LastName != "" {
		user.LastName = requestedUser.LastName
		query = query.Set("LAST_NAME", *requestedUser.LastName)
	}
	if requestedUser.Email != nil && *requestedUser.Email != "" {
		user.Email = requestedUser.Email
		query = query.Set("EMAIL", *requestedUser.Email)
	}
	if requestedUser.UserStatus != nil && *requestedUser.UserStatus != "" {
		user.UserStatus = requestedUser.UserStatus
		query = query.Set("USER_STATUS", *requestedUser.UserStatus)
	}

	hash, err := model.HashObject(user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	//Move id before we clear model
	requestedUser.Id = user.Id
	// clear the model
	user = &model.UserDTO{}
	query = query.Set("HASH", *hash).
		Where(squirrel.Eq{"ID": *requestedUser.Id}).
		Suffix("RETURNING ID, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT").
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
	query := squirrel.Delete("USERS").
		Where(squirrel.Eq{"HASH": id}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(db)
	_, err := query.Exec()
	if err != nil {
		return err
	}
	return nil
}
