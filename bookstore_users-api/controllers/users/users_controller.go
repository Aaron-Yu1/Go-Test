package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Aaron-Yu1/Go-Test/bookstore_users-api/domain/users"
	"github.com/Aaron-Yu1/Go-Test/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := errors.RestErr{
			Message: "invalid json body.",
			Status:  http.StatusBadRequest,
			Error:   "bad request.",
		}

		c.JSON(restErr.Status, restErr)

		fmt.Println(err)
		//TOD: retun bad requset to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}
