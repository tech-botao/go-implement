package repository

import "github.com/tech-botao/go-implement/cleanapp/app/domain/model"

// UserRepository 这个只是User那个的使用方法の接口定义
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(*model.User) error
}
