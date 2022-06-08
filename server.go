package main

import (
	"fmt"
	"log"
	"net"

	"github.com/j-ew-s/ms-curso-user-api/user"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":5001")

	if err != nil {
		log.Fatalf("Erro ao escutar a porta 5001 : %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := user.Server{}

	user.RegisterUserServiceServer(grpcServer, &userService)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(fmt.Sprintf("Erro ao levantar o gRPC Server porta 5001 : %v", err))
	}

}
