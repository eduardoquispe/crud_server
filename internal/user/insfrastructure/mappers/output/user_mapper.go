package mappers_output

import (
	"crud_server/internal/user/domain"
	"crud_server/internal/user/insfrastructure/dtos"
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

func ToUserResponse(user *domain.User) *dtos.UserResponse {
	return &dtos.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
}

func ToUserResponses(users []*domain.User) []*dtos.UserResponse {
	var userResponses []*dtos.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}