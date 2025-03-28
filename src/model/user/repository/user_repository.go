package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomanin model.UserDomanainInterface,
	) (model.UserDomanainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomanainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomanainInterface, *rest_err.RestErr)
	UpdateUser(id string, userDomanin model.UserDomanainInterface) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
	FindUserByEmailAndPassword(email string, password string) (model.UserDomanainInterface, *rest_err.RestErr)
}