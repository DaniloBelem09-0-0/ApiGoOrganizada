package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

var (
	JWT_SECRETE_KEY= "JWT_SECRETE_KEY"
)

func (ud *userDomanain) GenerreteTokeJwt() (string, *rest_err.RestErr) {
	secrete := os.Getenv(JWT_SECRETE_KEY)

	clains := jwt.MapClaims{
		"id": ud.id,
		"email": ud.email,
		"name": ud.name,
		"age": ud.age,
		"exp": time.Now().Add(time.Hour*1).Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	token_string, err := token.SignedString([]byte(secrete))
	if err != nil {
		return "" , rest_err.NewInternalServerError(fmt.Sprintf("Error trying to generate token, err=%s", err.Error()))
	}
	return token_string, nil
}

func VerifyToken(tokenValue string) (UserDomanainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRETE_KEY)
	token, err := jwt.Parse(RemmoveBearerPrefix(tokenValue), func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, fmt.Errorf("invalid token")
	})
	
	if err != nil {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	return &userDomanain{
		id: claims["id"].(string),
		email: claims["email"].(string),
		name: claims["name"].(string),
		age:  getIntClaim(claims, "age"),
	}, nil
}

func VerifyTokenMiddleWare(c *gin.Context) {
	secret := os.Getenv(JWT_SECRETE_KEY)
	tokenValue := RemmoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(RemmoveBearerPrefix(tokenValue), func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, fmt.Errorf("invalid token")
	})
	
	if err != nil {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(int(errRest.Code), errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(int(errRest.Code), errRest)
		c.Abort()
		return
	}

	userDomain := &userDomanain{
		id: claims["id"].(string),
		email: claims["email"].(string),
		name: claims["name"].(string),
		age:  getIntClaim(claims, "age"),
	}

	logger.Info(fmt.Sprintf("UserDomain: %v", userDomain), zap.String("journey", "verify token"))
	
}

func getIntClaim(claims jwt.MapClaims, key string) int8 {
    if val, ok := claims[key].(float64); ok { 
        return int8(val) 
    }
    return 0
}

func RemmoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer "){
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}