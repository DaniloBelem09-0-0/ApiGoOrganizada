package user

import (
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/service"
	"github.com/gin-gonic/gin"
)

func  NewUserControllerInterface(
	serviceInterface model.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface {
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	VerifyEmailHandler(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindUserById(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service model.UserDomainService
}