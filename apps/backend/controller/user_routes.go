package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	db "github.com/theCompanyDream/user-angular/apps/backend/repository"
)

func checkConstraints(c echo.Context) *db.User {
	user := &db.User{}
	if id := c.Param("id"); id != "" {
		i, _ := strconv.Atoi(id)
		user.Id = i
	}
	if userName := c.FormValue("user_name"); userName != "" {
		user.UserName = userName
	}
	if firstName := c.FormValue("firstName"); firstName != "" {
		user.FirstName = firstName
	}
	if lastName := c.FormValue("lastName"); lastName != "" {
		user.LastName = lastName
	}
	if email := c.FormValue("email"); email != "" {
		user.Email = email
	}
	if userStatus := c.FormValue("userStatus"); userStatus != "" {
		user.UserStatus = userStatus
	}
	if department := c.FormValue("department"); department != "" {
		user.Department = department
	}
	return user
}

func GetUser(c echo.Context) error {
	// Assume URL like /users/{id}
	// Extract the user ID from the URL and query the database
	request := checkConstraints(c)
	user, err := db.GetUser(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusFound, user)
}

func GetUsers(c echo.Context) error {
	// Assume URL like /users/{id}
	// Extract the user ID from the URL and query the database
	var page, limit int
	search := c.QueryParam("search")

	if queryLimit := c.QueryParam("limit"); queryLimit != "" {
		i, _ := strconv.Atoi(queryLimit)
		limit = i
	} else {
		limit = 25
	}
	if queryPage := c.QueryParam("page"); queryPage != "" {
		i, _ := strconv.Atoi(queryPage)
		limit = i
	} else {
		page = 1
	}

	users, error := db.GetUsers(search, page, limit)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusFound, users)
}

func CreateUser(c echo.Context) error {
	// Parse user details from the request body and insert into the database
	request := checkConstraints(c)
	user, error := db.CreateUser(request)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	// Parse user details from the request body and insert into the database
	request := checkConstraints(c)
	user, error := db.UpdateUser(*request)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	// Parse user details from the request body and insert into the database
	id := c.QueryParam("id")
	if id == "" {
		return errors.New("id must not be null")
	}
	idNumber, error := strconv.Atoi(id)
	if error != nil {
		return error
	}
	error = db.DeleteUser(idNumber)
	if error != nil {
		return error
	}
	return nil
}
