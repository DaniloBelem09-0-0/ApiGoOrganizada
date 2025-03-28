package user

import (
	"fmt"
	"net/http"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/validation"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/model/request"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomanainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
    zap.String("journey", "createUser"),
)

	var userRequest request.UserRequest
	
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while trying to marshal object", err,
		zap.String("journey", "createUser"), 
	)
		restErr := validation.ValidatorUserError(err)

		c.JSON(int(restErr.Code), restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password,
		userRequest.Name, 
	    userRequest.Age,
	)
	logger.Info(fmt.Sprintf("domain à ser salvo no banco de dados: %s", domain))
	domainResult, err := uc.service.CreateUserServices(domain)

	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}