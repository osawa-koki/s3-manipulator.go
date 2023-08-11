package app

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var (
	S3_BUCKET_NAME string
	S3_CLIENT      *s3.Client
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	S3_BUCKET_NAME = os.Getenv("S3_BUCKET_NAME")
	if S3_BUCKET_NAME == "" {
		log.Fatal("Error loading S3_BUCKET_NAME")
		os.Exit(1)
	}
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon S3 service client
	S3_CLIENT = s3.NewFromConfig(cfg)
}

func Handler() {
	listObject()
	putObject()
	getObject()
}
