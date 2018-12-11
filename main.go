package main

import (
	"challenge/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type server struct{}

func (s *server) CreateSkill(ctx context.Context, in *challenge.CreateSkillRequest) (*challenge.CreateSkillResponse, error) {
	return &challenge.CreateSkillResponse{Name: "asd", Id: 3}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	challenge.RegisterChallengeServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	fmt.Println("starting server on port " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
