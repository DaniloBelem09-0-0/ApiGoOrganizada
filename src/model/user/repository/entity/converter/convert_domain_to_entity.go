package converter

import (
	"fmt"

	model "github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/model/user/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomanainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Id: domain.GetId(),
		Email: domain.GetEmail(),
		Password: domain.GetPassword(),
		Name: domain.GetName(),
		Age: int8(domain.GetAge()),
	}
}

func ToString(e *entity.UserEntity) string {
	return fmt.Sprintf(
		"UserEntity{Email: %s, Name: %s, Age: %d, Password: %s, Id: %s}",
		e.Email,
		e.Name,
		e.Age,
		e.Password,
		e.Id,
	)
	// Nota: Password foi omitido por segurança
}