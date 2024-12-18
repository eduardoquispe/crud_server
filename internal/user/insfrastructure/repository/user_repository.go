package repository

import (
	"crud_server/internal/user/application/interfaces/output"
	"crud_server/internal/user/domain"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) output.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) FindAll() ([]*domain.User, error) {
	var users []*domain.User 
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindById(id int) (*domain.User, error) {
	var user domain.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) Save(user *domain.User) (*domain.User, error) {
	result := r.DB.Create(user)

	return user, result.Error
}

func (r *UserRepositoryImpl) Update(user *domain.User) (*domain.User, error) {
	fmt.Println(user)
	result := r.DB.Save(user)

	return user, result.Error
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&domain.User{}, id).Error
}
