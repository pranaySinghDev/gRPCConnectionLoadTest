package main

import (
	"log"
	"net"

	service "github.com/pranaySinghDev/gRPCConnectionLoadTest/user-app/service"

	user "github.com/pranaySinghDev/gRPCConnectionLoadTest/user-app/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	usr := service.Server{}
	user.RegisterUserServiceServer(grpcServer, &usr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
