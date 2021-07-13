package usecase

import (
	"github.com/google/uuid"
	"github.com/tech-botao/go-implement/cleanapp/app/domain/model"
	"github.com/tech-botao/go-implement/cleanapp/app/domain/repository"
	"github.com/tech-botao/go-implement/cleanapp/app/domain/service"
)

type UserUsecase interface{
	ListUser() ([]*User,  error)
	RegisterUser(email string) error
}

type userUseCase struct {
	repo repository.UserRepository
	service *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUseCase {
	return &userUseCase{
		repo:    repo,
		service: service,
	}
}

func (u *userUseCase) RegisterUser(email string) error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	if err := u.service.Duplicated(email); err != nil  {
		return err
	}

	user := model.NewUser(uid.String(), email)
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return nil
}

type (
	User struct {
		ID string
		Email string
	}
)

func toUser(users []*model.User) []*User {
	res  := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID: user.GetID(),
			Email: user.GetEmail(),
		}
	}
	return res
}