package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
)

// UploadFile - do upload a file to target server
// POST /upload
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	discussionID := c.PostForm("discussion_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// generate name based on current server date
	date := strings.Join(strings.Fields(time.Now().Format("2006.01.02 15:04:05")), "")
	file.Filename = "disc_" + discussionID + "_" + date + "_" + file.Filename

	log.Println(file.Filename)

	path := models.GoDotEnvVariable("PATH") + file.Filename
	errSaveMessage := c.SaveUploadedFile(file, path)
	if errSaveMessage != nil {
		c.JSON(http.StatusBadRequest, errSaveMessage.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"file_name": file.Filename})
}
