package persistence

import (
	"crud_server/internal/user/application/interfaces/output"
	"crud_server/internal/user/domain"
	mappers_output "crud_server/internal/user/insfrastructure/mappers/output"
	"crud_server/internal/user/insfrastructure/repository"
	"errors"
)

type UserPersistenceImpl struct {
	UserRepository repository.UserRepositoryImpl
}

func NewUserPersistence(userRepository repository.UserRepositoryImpl) output.UserPersistence {
	return &UserPersistenceImpl{
		UserRepository: userRepository,
	}
}

func (p *UserPersistenceImpl) Create(user *domain.User) (*domain.User, error) {
	entityUser := mappers_output.ToUserEntity(user)

	savedUser, err := p.UserRepository.Save(entityUser)
	if err != nil {
		return nil, err
	}

	domainUser := mappers_output.ToUserDomain(savedUser)
	return domainUser, nil
}

func (p *UserPersistenceImpl) FindById(id int) (*domain.User, error) {

	entityUser, err := p.UserRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if entityUser == nil {
		return nil, nil
	}

	return mappers_output.ToUserDomain(entityUser), nil
}	

func (p *UserPersistenceImpl) FindAll() ([]*domain.User, error) {
	entities, err := p.UserRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var users []*domain.User
	for _, entity := range entities {
		users = append(users, mappers_output.ToUserDomain(entity))
	}

	return users, nil
}

func (p *UserPersistenceImpl) Update(user *domain.User) (*domain.User, error) {
	entityUser := mappers_output.ToUserEntity(user)
	entityUser.ID = uint(user.Id)
	userFind, err := p.UserRepository.FindById(user.Id)
	if err != nil {
		return nil, err
	}

	if userFind == nil {
		return nil, errors.New("user not found")
	}

	savedUser, err := p.UserRepository.Update(entityUser)
	if err != nil {
		return nil, err
	}

	domainUser := mappers_output.ToUserDomain(savedUser)
	return domainUser, nil
}

func (p *UserPersistenceImpl) Delete(id int) error {
	return p.UserRepository.Delete(id)
}
