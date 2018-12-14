package main

import (
	"challenge/proto"
	"context"
	"database/sql"
	"fmt"
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

	rows, err := db.Query("SELECT * FROM skill")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows.Next()
	var id int
	var name string
	err = rows.Scan(&id, &name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &challenge.CreateSkillResponse{Name: name, Id: int32(id)}, nil
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
