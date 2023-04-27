package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Aaron-Yu1/Go-Test/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle json error
		return
	}

	result saveErr := services.CreateUser()

	fmt.Println(user)

	c.String(http.StatusNotImplemented, "implement me!")
}
