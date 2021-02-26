package main

import (
	"fmt"
	"github.com/ebina4yaka/go-grpc-tutorial/pokemon"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
	"net"
)

func main() {

	fmt.Println("Go gRPC Pokemon Server!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pokemonServer := pokemon.Server{}

	grpcServer := grpc.NewServer()

	pokemon.RegisterPokemonServiceServer(grpcServer, &pokemonServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
