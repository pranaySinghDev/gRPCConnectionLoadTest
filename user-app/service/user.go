package user

import (
	"context"

	user "github.com/pranaySinghDev/gRPCConnectionLoadTest/user-app/proto"
)

type Server struct {
}

func (s *Server) AddUser(ctx context.Context, us *user.User) (*user.User, error) {
	u := user.User{
		Id:   "one",
		Name: "Pranay",
		Age:  "27",
	}
	return &u, nil
}

func (s *Server) GetUser(ctx context.Context, us *user.User) (*user.User, error) {
	u := user.User{
		Id:   "one",
		Name: "Pranay",
		Age:  "27",
	}
	return &u, nil
}
