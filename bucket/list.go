package bucket

import (
    "fmt"
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "bucket"
)

func list() {
    fmt.Println("Listing buckets")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("eu-west-2")},
    )

    if err != nil {
        log.Fatal(err.Error())
    }

    service := s3.New(sess)

    result, err := service.ListBuckets(nil)
    if err != nil {
        log.Fatal("Error listing buckets")
    }

    for _, bucket := range result.Buckets {
        log.Printf("Bucket: %s\n", aws.StringValue(bucket.Name))
    }
}
