package view

import (
	"github.com/Joseeptessele/go-user-crud/src/controller/model/response"
	"github.com/Joseeptessele/go-user-crud/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
