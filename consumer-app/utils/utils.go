package utils

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	user "github.com/pranaySinghDev/gRPCConnectionLoadTest/user-app/proto"
)

var conn *grpc.ClientConn

func InitUserRPC() {
	connection, err := grpc.Dial("user-app:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	conn = connection
	// defer conn.Close()
}

func AddUser() *user.User {
	c := user.NewUserServiceClient(conn)
	response, err := c.AddUser(context.Background(), &user.User{Name: "Pranay", Age: "27"})
	if err != nil {
		log.Fatalf("Error when calling add user: %s", err)
	}
	return response

}

func GetUser() *user.User {
	c := user.NewUserServiceClient(conn)
	response, err := c.GetUser(context.Background(), &user.User{})
	if err != nil {
		log.Fatalf("Error when calling get user: %s", err)
	}
	return response
}
