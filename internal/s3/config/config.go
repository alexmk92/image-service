package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsClient struct {
    Session *session.Session
    Config *aws.Config
    S3Service *s3.S3
}

func (ac *AwsClient) Init(region string) error {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )

    if err != nil {
        return err
    }

    ac.Session = sess
    ac.S3Service = s3.New(sess)

    return nil
}


