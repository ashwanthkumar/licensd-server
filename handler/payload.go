package handler

import (
	"fmt"
	"log"
	"net/http"

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

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "/tmp/file")

	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"apiToken": apiToken,
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
