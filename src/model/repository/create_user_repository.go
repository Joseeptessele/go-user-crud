package repository

import (
	"context"
	"os"

	rest_err "github.com/Joseeptessele/go-user-crud/src/configuration"
	"github.com/Joseeptessele/go-user-crud/src/configuration/logger"
	"github.com/Joseeptessele/go-user-crud/src/model"
	"github.com/Joseeptessele/go-user-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGO_DB_USER_COLLECTION = "MONGO_DB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init create user repository", zap.String("journey", "createUser"))
	collection_name := os.Getenv(MONGO_DB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("error trying to create user", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	logger.Info("create user repository executed successfully",
		zap.String("journey", "createUser"),
		zap.String("userId", value.ID.Hex()))

	return converter.ConvertEntityToDomain(*value), nil
}
