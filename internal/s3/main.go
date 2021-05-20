package main

import (
	"github.com/alexmk92/image-service/internal/s3/bucket"
	s3 "github.com/alexmk92/image-service/internal/s3/config"
)

func main() {
    client := s3.AwsClient{}
    err := client.Init("us-west-2")
    if err != nil {
        panic(err)
    }

    b := bucket.NewBucketManager(&client)
    b.Create("test-buckets")
    b.List()
}
