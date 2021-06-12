package user

import (
	"fmt"
	"github.com/devmaufh/bookstore_users-api/domain/users"
	"github.com/devmaufh/bookstore_users-api/services"
	"github.com/devmaufh/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "List all users, not available yet.")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Get single user.")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	createdUser, storeError := services.CreateUser(user)
	if storeError != nil {
		c.JSON(storeError.StatusCode, storeError)
		return
	}
	c.JSON(http.StatusCreated, createdUser)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Search user")

}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Deletes user.")

}
