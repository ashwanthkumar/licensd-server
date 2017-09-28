package handler

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/ashwanthkumar/licensd-server/parser"
	"github.com/gin-gonic/gin"
)

// AddPayloadToDB handles the license file uploads from the client
func AddPayloadToDB(c *gin.Context) {
	// TODO(ashwanthkumar) Move this into a middleware to authenticate the requests that are flowing through
	apiToken := c.GetHeader("X-Licensd-API-Token")
	validateParam := missingParamInPayload(c)

	packageManager, exists := c.GetPostForm("package_manager")
	if !exists {
		validateParam("package_manager")
		return
	}
	log.Println("PackageManager=" + packageManager)
	buildVersion, exists := c.GetPostForm("version")
	if !exists {
		validateParam("version")
		return
	}
	log.Println("BuildVersion=" + buildVersion)
	buildMatrix, exists := c.GetPostForm("matrix")
	if !exists {
		validateParam("matrix")
		return
	}
	log.Println("BuildMatrix=" + buildMatrix)

	fileFormat, exists := c.GetPostForm("file_format")
	if !exists {
		fileFormat = parser.LICENSE_FINDER
		return
	}
	log.Println("FileFormat=" + fileFormat)

	buildURL, exists := c.GetPostForm("build_url")
	if !exists {
		validateParam("build_url")
		return
	}
	log.Println("BuildUrl=" + buildURL)

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
			"msg":    fmt.Sprintf("%s", err),
		})
		return
	}
	log.Println(file.Filename)
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
			"msg":    fmt.Sprintf("%s", err),
		})
		return
	}
	scanner := bufio.NewScanner(f)
	dependencies, err := parser.Parse(scanner, fileFormat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
			"msg":    fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"apiToken":     apiToken,
		"dependencies": dependencies,
	})
}

func missingParamInPayload(c *gin.Context) func(string) {
	return func(param string) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "FAILED",
			"msg":    fmt.Sprintf("%s is missing in the input payload", param),
		})
	}
}
