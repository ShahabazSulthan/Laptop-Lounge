package service

import (
	"Laptop_Lounge/pkg/config"
	"bytes"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func CreateSession(cfg *config.S3Bucket) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AccessKeyID,
			cfg.AccessKeySecret,
			"",
		),
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}

func UploadImageToS3(file *multipart.FileHeader, sess *session.Session) (string, error) {

	image, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer image.Close()

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	fmt.Println("11",uploader)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("bucket-laptop-lounge"),
		Key:    aws.String("product images/" + fileName),
		Body:   image,
		ACL:    aws.String("private"), // Change ACL to private
	})
	fmt.Println("err",upload)
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}


func UploadBytesToS3(data []byte, fileName string, folderName string, sess *session.Session) (string, error) {
	reader := bytes.NewReader(data)

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("bucket-laptop-lounge"),      // Update bucket name as needed
		Key:    aws.String(folderName + "/" + fileName), // Update folder structure if required
		Body:   reader,
		ACL:    aws.String("private"),
	})
	fmt.Println("errrrr",upload)
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}
