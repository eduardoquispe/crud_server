package mappers_input

import (
	"crud_server/internal/user/domain"
	"crud_server/internal/user/insfrastructure/dtos"
)

func ToUserDomain(user *dtos.UserRequest) *domain.User {
	return &domain.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
}
