package main

import (
	"github.com/tech-botao/go-implement/app/viper-cobra/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		panic(err)
	}
}
