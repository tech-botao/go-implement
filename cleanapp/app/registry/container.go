package registry

import (
	"github.com/sarulabs/di"
	"github.com/tech-botao/go-implement/cleanapp/app/domain/service"
	"github.com/tech-botao/go-implement/cleanapp/app/interface/persistence/memory"
	"github.com/tech-botao/go-implement/cleanapp/app/usecase"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	if err := builder.Add([]di.Def{
		{
			Name: "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{builder.Build()}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo :=  memory.NewUserRepository()
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}

