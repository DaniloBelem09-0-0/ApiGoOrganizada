package user

import (
	"net/http"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/validation"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/model/request"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init Login controller",
    zap.String("journey", "login"),
)

	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while trying to marshal object", err,
		zap.String("journey", "login"), 
	)
		restErr := validation.ValidatorUserError(err)

		c.JSON(int(restErr.Code), restErr)
		return
	}

	domain := model.NewUserDomainLogin(
		userRequest.Email, 
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)

	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User logged in successfully", zap.String("journey", "login"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}