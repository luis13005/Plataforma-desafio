package server

import (
	"PLATAFORMA-DESAFIO/application/usecases"
	"PLATAFORMA-DESAFIO/domain"
	"PLATAFORMA-DESAFIO/framework/pb"
	"context"
	"fmt"
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

func (UserServer *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	fmt.Println("ID:", req.GetUserId())
	oldPassword := req.GetOldPassword()
	UserServer.User.ID = req.GetUserId()
	UserServer.User.Name = req.GetName()
	UserServer.User.Email = req.GetEmail()
	UserServer.User.Password = req.GetNewPassword()
	UserServer.User.Token = req.GetToken()

	user, err := UserServer.UserUseCase.Update(&UserServer.User, oldPassword)
	if err != nil {
		log.Fatalf("Error during Updating user: %v", err)
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil
}
