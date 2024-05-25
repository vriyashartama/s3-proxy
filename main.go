package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	password := "(*ndias891kjnasdjni*8812"
	os.Setenv("PASSWORD", password)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[1:]

		session, _ := session.NewSession(&aws.Config{
			Region:           aws.String(os.Getenv("AWS_REGION")),
			Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT")),
			S3ForcePathStyle: aws.Bool(true),
			DisableSSL:       aws.Bool(true),
		})

		downloader := s3manager.NewDownloader(session)

		buf := aws.NewWriteAtBuffer([]byte{})

		numBytes, err := downloader.Download(buf,
			&s3.GetObjectInput{
				Bucket: aws.String(os.Getenv("AWS_BUCKET")),
				Key:    aws.String(filePath),
			},
		)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if numBytes > 0 {
			w.Write(buf.Bytes())
			return
		}
		w.Write([]byte("Error"))
	})

	http.ListenAndServe(":9990", nil)
}
