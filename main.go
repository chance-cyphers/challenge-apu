package main

import (
	"challenge/proto"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type server struct{}

func (s *server) CreateSkill(ctx context.Context, in *challenge.CreateSkillRequest) (*challenge.CreateSkillResponse, error) {
	connString, _ := DbConnectionString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var skillId int
	err = db.QueryRow(`INSERT INTO skill(name) VALUES ($1) RETURNING id`, in.Name).Scan(&skillId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &challenge.CreateSkillResponse{Name: in.Name, Id: int32(skillId)}, nil
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

	log.Println("starting server on port " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
