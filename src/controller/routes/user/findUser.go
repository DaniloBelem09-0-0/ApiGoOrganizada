package user

import (
	"net/http"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller", zap.String("journey", "findUserById"))

	userID := c.Param("userId")
	domainResult, err := uc.service.FindUserByIDServices(userID)
	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User found successfully", zap.String("journey", "findUserById"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context){
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "findByEmail"))
	
	userEmail := c.Param("userEmail")
	domainResult, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User found successfully", zap.String("journey", "findByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}