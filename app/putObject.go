package app

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getKeyString() string {
	date := time.Now()
	layout := "2006_01_02_15_04_05"
	return fmt.Sprintf("%s.txt", date.Format(layout))
}

func putObject() {
	// Create a new Amazon S3 bucket
	body := bytes.NewReader([]byte("body"))
	_, err := S3_CLIENT.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET_NAME),
		Key:    aws.String(getKeyString()),
		Body:   body,
	})
	if err != nil {
		log.Fatal(err)
	}
}
