package s3bucket

import (
	s3 "github.com/alexmk92/image-service/internal/s3/config"
	"github.com/aws/aws-sdk-go/aws"
	log "github.com/sirupsen/logrus"
)

type BucketManager struct {
    Client *s3.AwsClient
}

// Creates a new bucket that we can interact with
func NewBucketManager(s3Client *s3.AwsClient) *BucketManager {
    return &BucketManager{
        Client: s3Client,
    }
}

// List - print the contents of the bucket
func (b *BucketManager) List() {
    log.Info("Listing buckets")

    result, err := b.Client.ListBuckets(nil)
    if err != nil {
        log.Fatal("Error listing buckets")
    }

    for _, bucket := range result.Buckets {
        log.Printf("Bucket: %s\n", aws.StringValue(bucket.Name))
    }
}
