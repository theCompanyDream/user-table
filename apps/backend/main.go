package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/Masterminds/squirrel"
    _ "github.com/lib/pq"
)


var db *sql.DB
var sq squirrel.StatementBuilder

func initDB() {
    var err error
    db, err = sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    // Assume URL like /users/{id}
    // Extract the user ID from the URL and query the database
}

func createUser(w http.ResponseWriter, r *http.Request) {
    // Parse user details from the request body and insert into the database
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    // Parse user details from the request body and insert into the database
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    // Parse user details from the request body and insert into the database
}

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("/users", getUser)
    http.HandleFunc("/users/create", createUser)

    fmt.Println("Server is running...")
    http.ListenAndServe(":8080", nil)
}
