package users

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Aaron-Yu1/Go-Test/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	fmt.Print(string(bytes))
	c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
