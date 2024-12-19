package output

import "crud_server/internal/user/domain"

type UserPersistence interface {
	FindAll() ([]*domain.User, error)

	FindById(id int) (*domain.User, error)
	
	Create(user *domain.User) (*domain.User, error)

	Update(user *domain.User) (*domain.User, error)
		
	Delete(id int) error
}