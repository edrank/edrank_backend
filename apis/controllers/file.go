package controllers

import (
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func FileUploadController(c *gin.Context) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	bucket := utils.GetEnvWithKey("BUCKET_NAME")
	// region := utils.GetEnvWithKey("AWS_REGION")
	file, header, err := c.Request.FormFile("file")
	file_type := c.Request.FormValue("file_type")

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
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
		utils.SendError(c, http.StatusInternalServerError, errors.New("Failed to upload file : "+err.Error()))
		return
	}
	cid := c.GetInt("CollegeId")
	file_registry := models.FileRegistryModel{
		Cid:      cid,
		Location: fileKey,
		Url:      up.Location,
		Type:     file_type,
	}

	frId, err := models.CreateFileRegistry(file_registry)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, errors.New("File registry creation failed : "+err.Error()))
		return
	}

	utils.SendResponse(c, "File uploaded", map[string]any{
		"filepath":         up.Location,
		"file_registry_id": frId,
	})
}
