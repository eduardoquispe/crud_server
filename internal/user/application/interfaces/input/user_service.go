package input

import "crud_server/internal/user/domain"

type UserService interface {
	List() []*domain.User

	Read(id int) *domain.User
	
	Create(user *domain.User) (*domain.User, error)
	
	Update(id int, user *domain.User) (*domain.User, error)
	
	Delete(id int) error
}