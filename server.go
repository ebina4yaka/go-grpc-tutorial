package main

import (
	"fmt"
	"github.com/ebina4yaka/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	chatServer := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("faild to serve: %s", err)
	}
}
