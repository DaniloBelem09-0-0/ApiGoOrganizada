package routes

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:userId", user.FindUserById)
	r.GET("/usersByEmail/:userEmail", user.FindUserByEmail)
	r.POST("/user/", user.CreateUser)
	r.PUT("/user/:userId", user.UpdateUser)
	r.DELETE("/user/:userId", user.DeleteUser)

}