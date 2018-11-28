package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	sess, er := session.NewSession(&aws.Config{
//		Region: aws.String("us-east-1")},
//	)
//
//	if er != nil {
//		exitErrorf("error creating session", er)
//	}
//
//	svc := s3.New(sess)
//
//	s3Req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
//		Bucket: aws.String("challenge-video-bucket"),
//		Key:    aws.String("/test.txt"),
//	})
//
//	urlStr, er := s3Req.Presign(15 * time.Minute)
//	if er != nil {
//		exitErrorf("error presigning", er)
//	}
//
//	fmt.Fprintf(w, "{ \"url\": \"%s\"}", urlStr)
//}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello rob"}, nil
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
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
