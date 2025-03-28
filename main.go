package main

import (
	"context"
	"log"
	"path/filepath"

	dependencies "github.com/DaniloBelem09-0-0/ApiGoOrganizada/dependencies"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/database/mongodb"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           API Go Organizada
// @version         1.0
// @description     Esta é uma API organizada utilizando Gin e MongoDB
// @termsOfService  http://swagger.io/terms/

// @contact.name   Danilo Belem
// @contact.email  danilo@email.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
func main() {

	logger.Info("About to start aplication")
	path_dir := ""
	err := godotenv.Load(filepath.Join(path_dir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Erro trying to connect to database, error=%s \n",
			err.Error(),
		)
	}

	userController := dependencies.InitDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
		logger.Info("About to start aplication")

}

