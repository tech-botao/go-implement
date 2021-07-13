package v1

import (
	"context"
	"github.com/tech-botao/go-implement/cleanapp/app/interface/rpc/v1.0/protocol"
	"github.com/tech-botao/go-implement/cleanapp/app/usecase"
)

type UserService struct {
	userUsecase usecase.UserUsecase
}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {
	return
}

func NewUserService(userUsecase usecase.UserUsecase) *UserService {
	return &UserService{
		userUsecase: userUsecase,
	}
}

func (s *UserService) ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	users, err := s.userUsecase.ListUser()
	if err != nil {
		return nil, err
	}

	res := &protocol.ListUserResponseType{
		Users: toUser(users),
	}

	return res, nil
}

func toUser(users []*usecase.User) []*protocol.User {
	res := make([]*protocol.User, len(users))
	for i, user :=range users {
		res[i] = &protocol.User{
			Id : user.ID,
			Email: user.Email,
		}
	}
	return res
}

func (s *UserService) RegisterUser(ctx context.Context,in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error) {
	if err := s.userUsecase.RegisterUser(in.GetEmail()); err != nil {
		return nil, err
	}
	return &protocol.RegisterUserResponseType{}, nil
}





