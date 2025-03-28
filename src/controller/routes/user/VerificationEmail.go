package user

import (
	"net/http"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/service"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) VerifyEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	err := service.VerifyEmail(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
