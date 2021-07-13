package rpc

import (
	"github.com/tech-botao/go-implement/cleanapp/app/interface/rpc/v1.0"
	"github.com/tech-botao/go-implement/cleanapp/app/registry"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}