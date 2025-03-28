package routes

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/user"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/gin-gonic/gin"
	_ "github.com/DaniloBelem09-0-0/ApiGoOrganizada/docs" // Importa os docs gerados pelo Swag
	files "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.RouterGroup, 
	userController user.UserControllerInterface) {

	r.GET("/user/:userId", model.VerifyTokenMiddleWare, userController.FindUserById)
	r.GET("/usersByEmail/:userEmail", model.VerifyTokenMiddleWare, userController.FindUserByEmail)
	r.POST("/user/", model.VerifyTokenMiddleWare, userController.CreateUser)
	r.PUT("/user/:userId", model.VerifyTokenMiddleWare, userController.UpdateUser)
	r.DELETE("/user/:userId", model.VerifyTokenMiddleWare, userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
	r.GET("/verify_email", model.VerifyTokenMiddleWare, userController.VerifyEmailHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

}