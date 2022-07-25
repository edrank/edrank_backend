package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/edrank/edrank_backend/apis/utils"
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
