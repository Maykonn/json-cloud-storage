package src

import (
	"testing"
)

func TestNewS3Config(t *testing.T) {
	region := "us-east-1"
	bucket := "some-bucket"

	s3Config := NewS3Config(region, bucket)
	if s3Config.region != region || s3Config.bucket != bucket {
		t.Error("Invalid S3 configuration")
	}
}

func TestNewSession(t *testing.T) {
	region := "us-east-1"
	bucket := "some-bucket"

	s3Config := NewS3Config(region, bucket)

	_, err := NewSession(s3Config)
	if err != nil {
		t.Error(err)
	}
}
