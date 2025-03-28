package user

import (
	"fmt"
	"net/http"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
    zap.String("journey", "deleteUser"),)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(int(errRest.Code), errRest)
	}
	
	if errx := uc.service.DeleteUserServices(userId); errx != nil {
		logger.Error(
			"Error trying to call deleteUser from service", 
			fmt.Errorf("%v", errx),  // Convertendo errx para error, 
			zap.String("Journey", "deleteUser"),
		)
	}

	logger.Info("User delete successfully", zap.String("journey", "updateUser"))
	logger.Info("User delete successfully", zap.String("userId:", userId))
	
	c.Status(http.StatusOK)
}