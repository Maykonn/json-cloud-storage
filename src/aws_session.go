package src

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type s3Config struct {
	region string
	bucket string
}

func NewS3Config(region string, bucket string) *s3Config {
	return &s3Config{
		region: region,
		bucket: bucket,
	}
}

func NewSession(s3Config *s3Config) (*session.Session, error) {
	if s3Config.bucket == "" {
		panic("Bucket cannot be blank")
	}

	config := &aws.Config{Region: aws.String(s3Config.region)}
	return session.NewSession(config)
}
