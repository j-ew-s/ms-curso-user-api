package main

import (
	"log"
	"net"

	"github.com/j-ew-s/ms-curso-user-api/user"
	"google.google.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":5001")

	if err != nil {
		log.Fatalf("Erro ao escutar a porta 5001 : %v", err)
	}

	s := user.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Server(lis); err != nil {
		log.Fatal("Erro ao levantar o gRPC Server porta 5001 : %v", err)
	}

}
