package output

import "crud_server/internal/user/domain"

type UserRepository interface {
	FindAll() ([]*domain.User, error)

	FindById(id int) (*domain.User, error)
	
	Save(user *domain.User) (*domain.User, error)

	Update(user *domain.User) (*domain.User, error)
		
	Delete(id int) error
}