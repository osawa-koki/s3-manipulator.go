package app

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func listObject() {
	log.Println("===== ===== ===== ===== =====")
	log.Println("===== Listing Objects =====")
	log.Println("===== ===== ===== ===== =====")

	output, err := S3_CLIENT.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(S3_BUCKET_NAME),
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for i, object := range output.Contents {
		log.Printf("- %d: key=%s size=%d", i, aws.ToString(object.Key), object.Size)
	}
}
