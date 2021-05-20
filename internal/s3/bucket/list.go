package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	log "github.com/sirupsen/logrus"
    "github.com/alexmk92/image-service/internal/s3/config"
)

func list() {
    log.Info("Listing buckets")

    client := AwsClient{}
    err := client.Init("eu-west-2")
    if err != nil {
        panic(err)
    }

    result, err := client.S3Service.ListBuckets(nil)
    if err != nil {
        log.Fatal("Error listing buckets")
    }

    for _, bucket := range result.Buckets {
        log.Printf("Bucket: %s\n", aws.StringValue(bucket.Name))
    }
}
