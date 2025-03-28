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
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmailAndPassword(email string, password string) (model.UserDomanainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword repository")
	logger.Info(email)
	collectionName := os.Getenv(MONGO_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	var userEntity entity.UserEntity
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewForbiddenError("User not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user")
	} else {
		if userEntity.Password != password {
			return nil, rest_err.NewBadRequestError("Wrong Email or Password")
		}
	}

	return converter.ConvertEntityToDomain(userEntity), nil
}