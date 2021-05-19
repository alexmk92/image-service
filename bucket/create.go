package bucket

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func create() {
    fmt.Println("Creating buckets")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("eu-west-2")},
    )

    if err != nil {
        log.Fatal(err.Error())
    }

    service := s3.New(sess)
    input := &s3.CreateBucketInput{
        Bucket: aws.String("posterspy-main"),
    }

    result, err := service.CreateBucket(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
                case s3.ErrCodeBucketAlreadyExists:
                    log.Fatal("Bucket already exists")
                case s3.ErrCodeBucketAlreadyOwnedByYou:
                    log.Fatal("Bucket already owned by you")
                default:
                    log.Fatal(aerr.Error())
            }
        } else {
            log.Fatal(err.Error())
        }
    }

    log.Println(result)
}
