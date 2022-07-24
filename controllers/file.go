package controllers

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/edrank/edrank_backend/utils"
	"github.com/gin-gonic/gin"
)

func FileUploadController(c *gin.Context) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	bucket := utils.GetEnvWithKey("BUCKET_NAME")
	// region := utils.GetEnvWithKey("AWS_REGION")
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	filename := header.Filename
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String("edrank/" + filename),
		Body:   file,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file : " + err.Error(),
		})
		return
	}

	utils.SendResponse(c, "File uploaded", map[string]any{
		"filepath": up.Location,
	})
}
