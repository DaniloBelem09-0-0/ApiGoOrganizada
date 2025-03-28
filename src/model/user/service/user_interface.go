package service

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	// Métodos definidos na interface
	CreateUserServices(model.UserDomanainInterface) (model.UserDomanainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (model.UserDomanainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (model.UserDomanainInterface, *rest_err.RestErr)
	UpdateUserServices(id string, userDomanin model.UserDomanainInterface) *rest_err.RestErr
	DeleteUserServices(id string) *rest_err.RestErr
	LoginUserServices(userDomain model.UserDomanainInterface) (model.UserDomanainInterface, string, *rest_err.RestErr)
}
