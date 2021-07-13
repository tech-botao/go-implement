#!make
include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/8am/main.go

grpcc:
	grpcc --proto app/interface/rpc/v1.0/protocol/*.proto --address 127.0.0.1:${PORT} -i

protoc:
	#protoc --proto_path=. --go_out=plugins=grpc:./ app/interface/rpc/v1.0/protocol/*.proto
	protoc --go_out=. --go_opt=paths=import \
       --go-grpc_out=. --go-grpc_opt=paths=import \
       app/interface/rpc/v1.0/protocol/*.proto

install-grpcc:
	npm install -g grpcc