package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[1:]

		session, _ := session.NewSession(&aws.Config{
			Region:           aws.String(os.Getenv("AWS_REGION")),
			Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT")),
			S3ForcePathStyle: aws.Bool(true),
		})

		downloader := s3manager.NewDownloader(session)

		buf := aws.NewWriteAtBuffer([]byte{})

		numBytes, _ := downloader.Download(buf,
			&s3.GetObjectInput{
				Bucket: aws.String(os.Getenv("AWS_BUCKET")),
				Key:    aws.String(filePath),
			},
		)

		if numBytes > 0 {
			w.Write(buf.Bytes())
			return
		}
		w.Write([]byte("Error"))
	})

	http.ListenAndServe(":9990", nil)
}
