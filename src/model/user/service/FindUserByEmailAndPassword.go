package service

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
)

func (ud *userDomainService) FindUserByEmailAndPasswordServices(email string, password string) (model.UserDomanainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword model")

	userDomainRepository, err := ud.userRepository.FindUserByEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}