package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	//if er != nil {
	//	secret := os.Getenv("CLOUDCUBE_SECRET_ACCESS_KEY")
	//	id := os.Getenv("CLOUDCUBE_ACCESS_KEY_ID")
	//
	//	sess, er = session.NewSession(&aws.Config{
	//		Credentials: credentials.NewStaticCredentials(id, secret, "TOKEN"),
	//		Region:      aws.String("us-west-2")},
	//	)
	//}

	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
		fmt.Fprintf(w, "I love %s! ", aws.StringValue(b.Name))
	}
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
