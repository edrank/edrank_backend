package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
	"os"
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
