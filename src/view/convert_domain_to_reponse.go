package view

import (
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/controller/routes/model/response"
	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
)

func ConvertDomainToResponse(
	userDomain model.UserDomanainInterface,
) response.UserResponse {
	return response.UserResponse{
		Id: userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: int8(userDomain.GetAge()),
	}
}