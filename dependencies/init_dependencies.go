package dependencies

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependencies(
	database *mongo.Database,
) user.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return user.NewUserControllerInterface(service)
}