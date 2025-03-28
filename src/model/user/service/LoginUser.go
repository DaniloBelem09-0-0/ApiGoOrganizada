package service

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomanin model.UserDomanainInterface,
) (model.UserDomanainInterface, string, *rest_err.RestErr) {
	logger.Info("Init loginUser model", 
		zap.String("journey", "login"),
		zap.String("email", userDomanin.GetEmail()), // Logando o email do usuário
		zap.String("name", userDomanin.GetName()),   // Logando o nome do usuário
	)

	userDomanin.EncryptPassword()

	user, err := ud.FindUserByEmailAndPasswordServices(userDomanin.GetEmail(), userDomanin.GetPassword() )
	if err != nil {
		return nil, "", rest_err.NewBadRequestError(err.Err)
	}
	
	token, err := user.GenerreteTokeJwt()
	if err != nil {
		return nil, "", err
	}

	logger.Info("User logged in successfully",
		zap.String("journey", "login"),
		zap.String("email", userDomanin.GetEmail()),)

	return user, token, nil
}
