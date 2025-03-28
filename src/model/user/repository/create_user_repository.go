package repository

import (
	"context"
	"os"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGO_USER_COLLECTION = "MONGO_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomanin model.UserDomanainInterface,
) (model.UserDomanainInterface, *rest_err.RestErr) {
	logger.Info("Initi createUser repository[About to use Database]")

	collection_name := os.Getenv(MONGO_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	logger.Info("Using collection: " + collection_name)

	json := converter.ConvertDomainToEntity(userDomanin)
	logger.Info(converter.ToString(json));

	result , err := collection.InsertOne(context.Background(), json)

	if err != nil {
		return nil, rest_err.NewInternalServerError("Não foi criar usuário")
	} 

	if result.InsertedID == nil {
		return nil, rest_err.NewInternalServerError("Insertion not acknowledged by MongoDB")
	}

	json.Id = result.InsertedID.(primitive.ObjectID).Hex()

	if ur.databaseConnection == nil {
		return nil, rest_err.NewInternalServerError("Database connection is nil")
	}

	if result.InsertedID == nil {
		return nil, rest_err.NewInternalServerError("Insertion not acknowledged by MongoDB")
	}

	return converter.ConvertEntityToDomain(*json), nil
}