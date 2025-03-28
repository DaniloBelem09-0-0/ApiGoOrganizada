package service

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(id string, userDomanin model.UserDomanainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", 
		zap.String("journey", "createUser"),
		zap.String("email", userDomanin.GetEmail()), // Logando o email do usuário
		zap.String("name", userDomanin.GetName()),   // Logando o nome do usuário
	)

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
		zap.String("email", userDomanin.GetEmail()),
	)

	err := ud.userRepository.UpdateUser(id, userDomanin)
	if err != nil {
		return err
	}

	return nil
}
