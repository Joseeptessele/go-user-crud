package controller

import (
	"net/http"
	"net/mail"

	rest_err "github.com/Joseeptessele/go-user-crud/src/configuration"
	"github.com/Joseeptessele/go-user-crud/src/configuration/logger"
	"github.com/Joseeptessele/go-user-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById", zap.String("journey", "FindUserById"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userID", err, zap.String("journey", "FindUserById"))
		errorMessage := rest_err.NewBadRequestError("User ID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserById service", err, zap.String("journey", "FindUserById"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserById executed successfully", zap.String("journey", "FindUserById"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail", zap.String("journey", "FindUserByEmail"))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("userEmail is not a valid email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail service", err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByEmail executed successfully", zap.String("journey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
