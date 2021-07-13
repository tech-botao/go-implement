package main

import (
	"github.com/tech-botao/go-implement/cleanapp/app/interface/rpc"
	"github.com/tech-botao/go-implement/cleanapp/app/registry"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	server := grpc.NewServer()

	rpc.Apply(server, ctn)

	go func() {
		log.Printf("start grpc server port: %s", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("stopping grpc server...")
	server.GracefulStop()
	ctn.Clean()
}
