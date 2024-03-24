package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	db "github.com/theCompanyDream/user-angular/apps/backend/repository"
)

func checkConstraints(c echo.Context) db.User {
	user := db.User{}
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
	// Extract the user ID from the URL and query the database
	request := &db.User{}
	if id := c.Param("id"); id != "" {
		request.HashId = id
	}
	if userName := c.Param("user_name"); userName != "" {
		request.UserName = userName
	}
	if firstName := c.Param("firstName"); firstName != "" {
		request.FirstName = firstName
	}
	if lastName := c.Param("lastName"); lastName != "" {
		request.LastName = lastName
	}
	if email := c.Param("email"); email != "" {
		request.Email = email
	}
	user, err := db.GetUser(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusFound, user)
}

func GetUsers(c echo.Context) error {
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
	if id := c.Param("id"); id != "" {
		request.HashId = id
	}
	user, error := db.UpdateUser(request)
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
	err := db.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
