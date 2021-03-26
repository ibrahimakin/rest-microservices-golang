package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakin/rest-microservices-golang/bookstore_users-api/domain/users"
	"github.com/ibrahimakin/rest-microservices-golang/bookstore_users-api/services"
	"github.com/ibrahimakin/rest-microservices-golang/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error when read bytes")
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError("invalid JSON Body")
		// TODO: Return bad request to the caller.
		c.JSON(restErr.Status, restErr)
		return
	}
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	fmt.Println(err.Error())
	// 	// TODO: Handle json error
	// 	c.String(http.StatusInternalServerError, "Error when unmarshal bytes")
	// 	return
	// }
	fmt.Println("User", user)
	fmt.Println("Bytes", string(bytes))

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		// TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewNotFoundError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		// TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusOK, user)
}
