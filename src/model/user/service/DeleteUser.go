package service

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(
	id string,
) (*rest_err.RestErr) {
	logger.Info("Init createUser model", 
		zap.String("journey", "deleteUser"),
	)

	
	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
