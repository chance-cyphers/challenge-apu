package main

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	message, err := getBuckets()

	if err != nil {
		fmt.Println(err.Error())
	}

	return &pb.HelloReply{Message: message}, err
}

func getBuckets() (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}

	projectID := "challenge-api"
	it := client.Buckets(ctx, projectID)
	buffer := bytes.Buffer{}
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		_, _ = buffer.WriteString(battrs.Name)
		_, _ = buffer.WriteString("\n")
	}
	return buffer.String(), nil
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
	pb.RegisterGreeterServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	fmt.Println("starting server on port " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
