package controller

import (
	"net/http"

	"github.com/Joseeptessele/go-user-crud/src/configuration/logger"
	"github.com/Joseeptessele/go-user-crud/src/configuration/validation"
	"github.com/Joseeptessele/go-user-crud/src/controller/model/request"
	"github.com/Joseeptessele/go-user-crud/src/model"
	"github.com/Joseeptessele/go-user-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err, zap.String("journey", "createUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("error trying to call create user service", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
