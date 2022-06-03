package main

import (
	"log"
	"net"

	"google.google.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":5001")

	if err != nil {
		log.Fatalf("Erro ao escutar a porta 5001 : %v", err)
	}

	grpcServer := grpc.NewServer()

}
