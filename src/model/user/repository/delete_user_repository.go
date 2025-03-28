package repository

import (
	"context"
	"os"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	id string,
) *rest_err.RestErr {
	logger.Info("Initi deleteUser repository[About to use Database]")

	collection_name := os.Getenv(MONGO_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	logger.Info("Using collection: " + collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rest_err.NewBadRequestError("UserID is not a valid hex")
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		logger.Error("Erro trying to delete", err, zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUSer repository executed successfully",
		zap.String("userID", id),
		zap.String("journey", "deleteUser"),
	)

	return nil
}