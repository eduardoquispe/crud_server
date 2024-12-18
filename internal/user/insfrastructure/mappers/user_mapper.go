package mappers

import (
	"crud_server/internal/user/domain"
	"crud_server/internal/user/insfrastructure/entity"

	"gorm.io/gorm"
)

func ToUserDomain(entity *entity.User) *domain.User { 
	return &domain.User{
		Id:        int(entity.ID),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Age:       entity.Age,
	}
}

func ToUserEntity(user *domain.User) *entity.User {
	return  &entity.User{
		Model:     gorm.Model{ID: uint(user.Id)},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
}