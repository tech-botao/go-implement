package v1

import (
	"github.com/tech-botao/go-implement/cleanapp/app/interface/rpc/v1.0/protocol"
	"github.com/tech-botao/go-implement/cleanapp/app/registry"
	"github.com/tech-botao/go-implement/cleanapp/app/usecase"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	// 这个不好使, protobuf 转换程序有点儿问题
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}

