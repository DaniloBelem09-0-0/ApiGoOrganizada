package repository

import (
	"context"
	"os"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository/entity"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByID(id string) (model.UserDomanainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID repository")
	logger.Info(id)
	collectionName := os.Getenv(MONGO_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid user ID")
	}

	var userEntity entity.UserEntity
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}, ).Decode(&userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewBadRequestError("User not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user")
	}

	return converter.ConvertEntityToDomain(userEntity), nil
}