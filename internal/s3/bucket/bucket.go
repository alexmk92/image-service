package s3bucket

import (
	s3Local "github.com/alexmk92/image-service/internal/s3/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

type BucketManager struct {
    Client *s3Local.AwsClient
}

// Creates a new bucket that we can interact with
func NewBucketManager(s3Client *s3Local.AwsClient) *BucketManager {
    return &BucketManager{
        Client: s3Client,
    }
}

// List - print the contents of the bucket
func (b *BucketManager) List() {
    log.Info("Listing buckets")

    result, err := b.Client.S3Service.ListBuckets(nil)
    if err != nil {
        log.Fatal("Error listing buckets")
    }

    for _, bucket := range result.Buckets {
        log.Printf("Bucket: %s\n", aws.StringValue(bucket.Name))
    }
}

// Create - creates a new bucket
func (b *BucketManager) Create(bucketName string) {
    input := &s3.CreateBucketInput{
        Bucket: aws.String(bucketName),
    }

    result, err := b.Client.S3Service.CreateBucket(input)
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

    log.Info(result)
}
