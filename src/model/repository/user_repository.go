package repository

import (
	rest_err "github.com/Joseeptessele/go-user-crud/src/configuration"
	"github.com/Joseeptessele/go-user-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
