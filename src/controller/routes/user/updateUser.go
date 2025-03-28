package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/validation"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/model/request"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
    zap.String("journey", "updateUser"),)

	var userRequest request.UserRequestUpdate
	userId := c.Param("userId")
	
	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error while trying to marshal object", err,
		zap.String("journey", "updateUser"), 
	)
		restErr := validation.ValidatorUserError(err)

		c.JSON(int(restErr.Code), restErr)
		return
	}

	domain := model.NewUserDomainUpdate(
		userRequest.Name, 
	    userRequest.Age,
	)

	logger.Info(fmt.Sprintf("domain à ser salvo no banco de dados: %s", domain))
	err := uc.service.UpdateUserServices(userId, domain)

	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "updateUser"))
	logger.Info("User created successfully", zap.String("userId:", userId))
	
	c.Status(http.StatusOK)
}