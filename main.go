package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/cors"
)

// func uploadToServer(r io.Reader, filename string) (string, error) {
// 	dst, err := os.Create("./" + filename)
// 	defer dst.Close()
// 	if err != nil {
// 		return "", err
// 	}
// 	//copy the uploaded file to the destination file
// 	if _, err := io.Copy(dst, r); err != nil {
// 		return "", nil
// 	}
// 	return filename, nil
// }

var (
	S3_REGION = getEnv("REGION", "eu-west-1")
	S3_BUCKET = os.Getenv("BUCKET")
	ENDPOINT  = getEnv("ENDPOINT", "")
	// AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	// AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	S3_DIR      = getEnv("DIRECTORY", "")
	PORT        = getEnv("PORT", "80")
	S3_BASE_URL = fmt.Sprintf("https://%v.s3-%v.amazonaws.com/", S3_BUCKET, S3_REGION)
	// # S3_ENDPOINT = "https://fra1.digitaloceanspaces.com"
	// # S3_BASE_URL = "https://instagrammedias.fra1.cdn.digitaloceanspaces.com/"
	// # S3_BASE_URL = "https://storage.googleapis.com"
	// # gcp S3_BASE_URL = f"https://storage.googleapis.com/{S3_BUCKET}/"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

var creds *credentials.Credentials = credentials.NewEnvCredentials()
var awsConfig aws.Config = aws.Config{Region: aws.String(S3_REGION), Credentials: creds, Endpoint: aws.String(ENDPOINT)}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func uploadToS3(r io.Reader, filename string) (string, error) {

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&awsConfig))

	// Create an uploader with the session and default options
	//uploader := s3manager.NewUploader(sess)

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = 2            // default is 5
	})

	//defer f.Close()
	hex, err := randomHex(20)
	if err != nil {
		return "", err
	}
	key := S3_DIR + hex + "_" + filename
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(key),
		Body:   r,
		ACL:    aws.String("public-read"),
	})

	//in case it fails to upload
	if err != nil {
		log.Printf("failed to upload file, %v", err)
		return "", err
	}
	log.Printf("file uploaded to %s\n", result.Location)
	return result.Location, nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case "GET":
	// 	display(w, "upload", nil)
	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//get the multipart reader for the request.
		reader, err := r.MultipartReader()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//copy each part to destination.
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}

			//if part.FileName() is empty, skip this iteration.
			if part.FileName() == "" {
				continue
			}
			url, err := uploadToS3(part, part.FileName())

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			io.WriteString(w, url)
			return
		}
		//display success message.
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)
	mux.HandleFunc("/upload", handle)
	log.Println("starting listening")
	err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), handler)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
