package app

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func generateRandom(i1, i2 int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return r.Intn(i2-i1) + i1
}

func deleteObject() {
	log.Println("===== ===== ===== ===== =====")
	log.Println("===== Deleting Objects =====")
	log.Println("===== ===== ===== ===== =====")

	output, err := S3_CLIENT.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(S3_BUCKET_NAME),
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// 50%の確率で削除する。
	if generateRandom(0, 1) != 0 {
		log.Println("Skip deleting objects...")
		return
	}

	object := output.Contents[generateRandom(0, len(output.Contents))]
	log.Println("Deleting object: ", aws.ToString(object.Key))

	_, err = S3_CLIENT.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(S3_BUCKET_NAME),
		Key:    aws.String(aws.ToString(object.Key)),
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
