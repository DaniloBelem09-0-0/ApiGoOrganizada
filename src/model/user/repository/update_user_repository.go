package repository

import (
	"context"
	"os"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(id string, userDomanin model.UserDomanainInterface) *rest_err.RestErr {
	logger.Info("Initi updateUser repository[About to use Database]")

	collection_name := os.Getenv(MONGO_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	logger.Info("Using collection: " + collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rest_err.NewBadRequestError("UserID is not a valid hex")
	}

	value := converter.ConvertDomainToEntity(userDomanin)
	logger.Info(converter.ToString(value));

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		logger.Error("Erro trying to updateUser", err, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUSer repository executed successfully",
		zap.String("userID", value.Id),
		zap.String("journey", "updateUser"),
	)

	return nil
}