package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func OnBoardCollegeController(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./tmp/" + file.Filename
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)

	f, err := excelize.OpenFile(dst)

	if err != nil {
		log.Fatal(err)
	}
	cols, err := f.GetRows("Sheet1")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cols)

	c1, err := f.GetCellValue("Sheet1", "A1")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c1)

	c2, err := f.GetCellValue("Sheet1", "A4")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)

	c3, err := f.GetCellValue("Sheet1", "B2")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c3)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
