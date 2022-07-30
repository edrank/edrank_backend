package services

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws() (*session.Session, error) {
	AccessKeyID = utils.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = utils.GetEnvWithKey("AWS_ACCESS_KEY_SECRET")
	MyRegion = utils.GetEnvWithKey("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func DownloadFromS3(c *gin.Context, key string) error {
	sess := c.MustGet("sess").(*session.Session)
	bucket := utils.GetEnvWithKey("BUCKET_NAME")
	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create("tmp/" + key)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

	if err != nil {
		return err
	}

	return nil
}

func UploadToS3(c *gin.Context, upload_type string, path string, filename string) (string, error) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	bucket := utils.GetEnvWithKey("BUCKET_NAME")
	// region := utils.GetEnvWithKey("AWS_REGION")

	if upload_type == "request" {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return "", err
		}

		filename := header.Filename
		fileKey := "edrank/" + filename
		//upload to the s3 bucket
		up, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			ACL:    aws.String("public-read"),
			Key:    aws.String(fileKey),
			Body:   file,
		})
		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return "", err
		}

		return up.Location, nil
	}
	file, err := os.ReadFile(path)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return "", err
	}

	fileKey := "edrank/" + filename

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileKey),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return "", err
	}

	return up.Location, nil

}
