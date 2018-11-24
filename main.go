package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sess, er := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if er != nil {
		exitErrorf("error creating session", er)
	}

	svc := s3.New(sess)

	s3Req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("challenge-video-bucket"),
		Key:    aws.String("/test.txt"),
	})

	urlStr, er := s3Req.Presign(15 * time.Minute)
	if er != nil {
		exitErrorf("error presigning", er)
	}

	fmt.Fprintf(w, "{ \"url\": \"%s\"}", urlStr)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
