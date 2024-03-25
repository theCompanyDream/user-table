package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	model "github.com/theCompanyDream/user-angular/apps/backend/models"
	db "github.com/theCompanyDream/user-angular/apps/backend/repository"
)

func checkConstraints(c echo.Context) model.User {
	user := model.User{}
	if userName := c.FormValue("user_name"); userName != "" {
		user.UserName = &userName
	}
	if firstName := c.FormValue("first_name"); firstName != "" {
		user.FirstName = &firstName
	}
	if lastName := c.FormValue("last_name"); lastName != "" {
		user.LastName = &lastName
	}
	if email := c.FormValue("email"); email != "" {
		user.Email = &email
	}
	if userStatus := c.FormValue("user_status"); userStatus != "" {
		user.UserStatus = &userStatus
	}
	if department := c.FormValue("department"); department != "" {
		user.Department = &department
	}
	return user
}

// GetUser godoc
// @Summary Get a single user
// @Description Get a user by their ID or username
// @Tags user
// @Accept json
// @Produce json
// @Param id path string false "User ID"
// @Param user_name path string false "Username"
// @Success 302 {object} models.User "User Found"
// @Failure 400 {object} object "Bad Request"
// @Router /user/{id} [get]
func GetUser(c echo.Context) error {
	// Extract the user ID from the URL and query the database
	request := &model.User{}
	if id := c.Param("id"); id != "" {
		request.HashId = &id
	}
	if userName := c.Param("user_name"); userName != "" {
		request.UserName = &userName
	}
	if firstName := c.Param("first_name"); firstName != "" {
		request.FirstName = &firstName
	}
	if lastName := c.Param("last_name"); lastName != "" {
		request.LastName = &lastName
	}
	if email := c.Param("email"); email != "" {
		request.Email = &email
	}
	err := c.Validate(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	user, err := db.GetUser(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusFound, user)
}

// GetUsers godoc
// @Summary Get multiple users
// @Description Get a list of users, with optional search, pagination, and limit
// @Tags user
// @Accept json
// @Produce json
// @Param search query string false "Search Term"
// @Param limit query int false "Limit"
// @Param page query int false "Page Number"
// @Success 302 {object} []models.User "Users Found"
// @Failure 400 {object} object "Bad Request"
// @Router /users [get]
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

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user with the provided information
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "User object"
// @Success 201 {object} models.User "User Created"
// @Failure 400 {object} object "Bad Request"
// @Router /user [post]
func CreateUser(c echo.Context) error {
	// Parse user details from the request body and insert into the database
	request := checkConstraints(c)
	err := c.Validate(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	user, error := db.CreateUser(request)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user's information by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body User true "User object"
// @Success 200 {object} models.User "User Updated"
// @Failure 400 {object} object "Bad Request"
// @Router /user/{id} [put]
func UpdateUser(c echo.Context) error {
	// Parse user details from the request body and insert into the database
	request := checkConstraints(c)
	err := c.Validate(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	if id := c.Param("id"); id != "" {
		request.HashId = &id
	}
	user, error := db.UpdateUser(request)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "User Deleted"
// @Failure 400 {object} object "Bad Request"
// @Router /user/{id} [delete]
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
