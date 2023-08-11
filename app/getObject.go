package app

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getObject() {
	log.Printf("===== ===== ===== ===== =====")
	log.Printf("===== Getting Object =====")
	log.Printf("===== ===== ===== ===== =====")

	output, err := S3_CLIENT.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(S3_BUCKET_NAME),
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for i, object := range output.Contents {
		// ファイルの中身を取得する
		output, err := S3_CLIENT.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(S3_BUCKET_NAME),
			Key:    aws.String(aws.ToString(object.Key)),
		})
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(output.Body)
		log.Printf("- %d: key=%s size=%d body=%s", i, aws.ToString(object.Key), object.Size, buf.String())
	}
}
