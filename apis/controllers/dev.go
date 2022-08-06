package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetOnBoardingFileController(c *gin.Context) {
	devAuth := c.GetHeader("x-edrank-dev-auth")
	if devAuth != utils.GetEnvWithKey("DEV_TOKEN") {
		utils.SendError(c, http.StatusUnauthorized, errors.New("Invalid Dev Token."))
		return
	}

	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	bucket := utils.GetEnvWithKey("BUCKET_NAME")
	// region := utils.GetEnvWithKey("AWS_REGION")
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// get the file extentsion
	fileE := strings.Split(header.Filename, ".")
	fileExt := fileE[len(fileE)-1]

	fmt.Println(header, fileE, fileExt)

	if fileExt != "xlsx" {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid file format. Please upload only .xlsx file"))
		return
	}

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String("edrank/constants/EDRANK_ONBOARDING.xlsx"),
		Body:   file,
	})
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Failed to upload file : "+err.Error()))
		return
	}

	utils.SendResponse(c, "File set", map[string]any{
		"filepath": up.Location,
	})
}

func GenPasswordHash(c *gin.Context) {
	var body struct {
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	utils.SendResponse(c, "", map[string]any{
		"a": string(bytes),
	})
}
