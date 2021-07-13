package service

import (
	"fmt"
	"github.com/tech-botao/go-implement/cleanapp/app/domain/repository"
)

// UserService 包含repository的接口, 一部分业务逻辑可以写这里
type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Duplicated(email string) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}
	if user != nil {
		return fmt.Errorf("%s already exists",  email)
	}
	return nil
}
