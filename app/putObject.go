package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	prefixes = []string{
		"cute",
		"cool",
		"beautiful",
		"pretty",
		"lovely",
		"adorable",
		"gorgeous",
		"charming",
		"attractive",
		"elegant",
		"exquisite",
		"magnificent",
		"delicate",
		"graceful",
		"pleasing",
	}
	animals = []string{
		"cat",
		"dog",
		"rabbit",
		"hamster",
		"goldfish",
		"turtle",
		"bird",
		"snake",
		"fish",
		"mouse",
		"pig",
		"cow",
		"sheep",
		"chicken",
		"horse",
		"duck",
	}
)

func getBodyMessageString() string {
	return fmt.Sprintf("%s %s", prefixes[time.Now().Unix()%int64(len(prefixes))], animals[time.Now().Unix()%int64(len(animals))])
}

func getKeyString() string {
	date := time.Now()
	layout := "2006_01_02_15_04_05"
	return fmt.Sprintf("%s.txt", date.Format(layout))
}

type Body struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func putObject() {
	log.Println("===== ===== ===== ===== =====")
	log.Println("===== Putting Object =====")
	log.Println("===== ===== ===== ===== =====")

	bodyMessage := getBodyMessageString()
	body := Body{
		Message:   bodyMessage,
		Timestamp: time.Now(),
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = S3_CLIENT.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET_NAME),
		Key:    aws.String(getKeyString()),
		Body:   bytes.NewReader(jsonBody),
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
