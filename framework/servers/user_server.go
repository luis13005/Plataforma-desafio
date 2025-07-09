package server

import (
	"PLATAFORMA-DESAFIO/application/usecases"
	"PLATAFORMA-DESAFIO/domain"
	"PLATAFORMA-DESAFIO/framework/pb"
	"context"
	"log"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	UserServer.User.Name = req.GetName()
	UserServer.User.Email = req.GetEmail()
	UserServer.User.Password = req.GetPassword()

	user, err := UserServer.UserUseCase.Create(&UserServer.User)
	if err != nil {
		log.Fatalf("Error during Creating user: %v", err)
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil
}
