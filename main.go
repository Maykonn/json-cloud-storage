package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"maykonn/json-cloud-storage/src"
	"os"
)

func main() {
	if _, err := (godotenv.Load()).(error); err {
		panic("Error when loading .env file")
	}

	fmt.Println(os.Getenv("S3_BUCKET"))
	s3Config := src.NewS3Config(os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET"))
	awsSession := src.NewSession(s3Config)

	err := src.Save(awsSession, s3Config, "body.json")
	if err != nil {
		panic(err)
	}
}
