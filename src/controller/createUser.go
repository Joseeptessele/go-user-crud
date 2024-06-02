package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joseeptessele/go-user-crud/src/configuration/validation"
	"github.com/Joseeptessele/go-user-crud/src/controller/model/request"
	"github.com/Joseeptessele/go-user-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusCreated, response)
}
