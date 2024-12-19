package repository

import (
	"crud_server/internal/user/insfrastructure/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) FindAll() ([]*entity.User, error) {
	var users []*entity.User 
	err := r.DB.Order("id asc").Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) Save(user *entity.User) (*entity.User, error) {
	result := r.DB.Create(user)

	return user, result.Error
}

func (r *UserRepositoryImpl) Update(user *entity.User) (*entity.User, error) {
	result := r.DB.Save(user)

	return user, result.Error
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&entity.User{}, id).Error
}
