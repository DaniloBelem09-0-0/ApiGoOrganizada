package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"go.uber.org/zap"
)

var verificationTokens = make(map[string]string)
var verifiedEmails = make(map[string]bool)
var mu sync.Mutex

func (ud *userDomainService) CreateUserServices(
	userDomanin model.UserDomanainInterface,
) (model.UserDomanainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", 
		zap.String("journey", "createUser"),
		zap.String("email", userDomanin.GetEmail()),
		zap.String("name", userDomanin.GetName()),   
	)

	user, _ := ud.FindUserByEmailServices(userDomanin.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already in use by another account")
	}
	
	token := generateVerificationToken()
	storeVerificationToken(userDomanin.GetEmail(), token)
	verificationLink := fmt.Sprintf("http://localhost:8080/verify_email?token=%s", token)

	if err := SendVerificationEmail(userDomanin.GetEmail(), verificationLink); err != nil {
		return nil, rest_err.NewInternalServerError("Failed to send verification email")
	}

	logger.Info("Waiting for email verification", zap.String("email", userDomanin.GetEmail()))

	for range 300 {
		mu.Lock()
		if verifiedEmails[userDomanin.GetEmail()] {

			mu.Unlock()
			break
		}
		mu.Unlock()
		time.Sleep(2 * time.Second) 
	}

	userDomanin.EncryptPassword()
	
	logger.Info("User verified and created successfully",
		zap.String("journey", "createUser"),
		zap.String("email", userDomanin.GetEmail()),)

	userDomainRepository, errx := ud.userRepository.CreateUser(userDomanin)
	if errx != nil {
		return nil, errx
	}

	return userDomainRepository, nil
}