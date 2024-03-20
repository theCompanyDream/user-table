package controller

import (
	"github.com/labstack/echo/v4"
    "strconv"
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

func GetUser(c echo.Context) (*db.User, error)  {
    // Assume URL like /users/{id}
    // Extract the user ID from the URL and query the database
    request := checkConstraints(c)
    user, err := db.GetUser(request)
    if err != nil {
        return nil, c.JSON(http.)
    }
    return c.Json(http.StatusOK, user), nil
}

func GetUsers(c echo.Context) (db.User, error)  {
    // Assume URL like /users/{id}
    // Extract the user ID from the URL and query the database
    search := c.QueryParam("search")
    page := c.QueryParam("page")
    limit := c.QueryParam("limit")

    if limit == "" {
        limit = 30
    }
    if page == "" {
        page = 1
    }

    user := db.GetUsers(search, page)
    return user, nil
}

func CreateUser(c echo.Context) (db.User, error) {
    // Parse user details from the request body and insert into the database
    request := checkConstraints(c)
    user := db.CreateUser(request)
    return user, nil
}

func UpdateUser(c echo.Context) (db.User, error) {
    // Parse user details from the request body and insert into the database
    request := checkConstraints(c)
    user := db.UpdateUser(request)
    return user, nil
}

func DeleteUser(c echo.Context) (error) {
    // Parse user details from the request body and insert into the database
    id := c.QueryParam("id")
    if id != nil {
        return errror.New("id must not be null")
    }
    error := db.Delete(request)
    if error != nil {
        return error
    }
    return nil
}