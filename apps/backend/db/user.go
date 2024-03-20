package model

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"os"
	"fmt"
	"log"
)

type User struct {
    ID    int    `json:"id"`
    FName  string `json:"first_name"`
	lName  string `json:"last_name"`
    Email string `json:"email"`
	user_status string `json:"user_status"`
	department string `json:"department"`
}

var db *sql.DB
var sq squirrel.StatementBuilder

func initDB() {
    var err error

	connectStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))

    db, err = sql.Open("postgres", connectStr)
    if err != nil {
        log.Fatal(err)
    }

	sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db)
}

func getUser(c User) {
    // Assume URL like /users/{id}
    // Extract the user ID from the URL and query the database

}

func createUser(c User) {
    // Parse user details from the request body and insert into the database
}

func updateUser(c User) {
    // Parse user details from the request body and insert into the database
}

func deleteUser(c User) {
    // Parse user details from the request body and insert into the database
}