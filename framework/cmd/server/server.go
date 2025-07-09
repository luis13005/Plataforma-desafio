package main

import (
	"PLATAFORMA-DESAFIO/application/repositories"
	"PLATAFORMA-DESAFIO/application/usecases"
	"PLATAFORMA-DESAFIO/framework/pb"
	server "PLATAFORMA-DESAFIO/framework/servers"
	"PLATAFORMA-DESAFIO/framework/utils"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func main() {

	db = utils.ConectDB()
	db.LogMode(true)

	port := flag.Int("port", 5050, "Choose the server port")
	flag.Parse()
	log.Printf("Server started on port: %v", *port)

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

func setUpUserServer() *server.UserServer {
	userRepository := repositories.UserRepositoryDb{DB: db}
	userServer := server.NewUserServer()
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}

	return userServer
}
