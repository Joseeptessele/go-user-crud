package repository

import (
	"context"
	"os"

	rest_err "github.com/Joseeptessele/go-user-crud/src/configuration"
	"github.com/Joseeptessele/go-user-crud/src/configuration/logger"
	"github.com/Joseeptessele/go-user-crud/src/model"
	"github.com/Joseeptessele/go-user-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGO_DB_USER_COLLECTION = "MONGO_DB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init create user repository")
	collection_name := os.Getenv(MONGO_DB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
