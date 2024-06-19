package service

import (
	rest_err "github.com/Joseeptessele/go-user-crud/src/configuration"
	"github.com/Joseeptessele/go-user-crud/src/configuration/logger"
	"github.com/Joseeptessele/go-user-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById", zap.String("journey", "FindUserById"))
	return ud.userRepository.FindUserById(id)

}
func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail", zap.String("journey", "FindUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}
