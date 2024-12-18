package services

import (
	"crud_server/internal/user/application/interfaces/input"
	"crud_server/internal/user/application/interfaces/output"
	"crud_server/internal/user/domain"
)

type UserServiceImpl struct {
	repo output.UserRepository
}

func NewUserServiceImpl(repo output.UserRepository) input.UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) List() []*domain.User {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil
	}
	return users
}

func (s *UserServiceImpl) Read(id int) *domain.User {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil
	}
	return user
}

func (s *UserServiceImpl) Create(user *domain.User) (*domain.User, error) {
	user, err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) Update(id int, user *domain.User) (*domain.User, error) {
	user, err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}
