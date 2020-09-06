package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/controller"
	"github.com/zarszz/discussion-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// user context
	r.GET("/users", controller.FindUsers)
	r.POST("/user", controller.InsertUser)

	// category context
	r.GET("/categories", controller.FindCategories)
	r.GET("/category/:id", controller.FindCategory)
	r.POST("/category", controller.InsertCategory)
	r.PUT("/category/:id", controller.UpdateCategory)
	r.DELETE("/category/:id", controller.DeleteCategory)

	// discussion context
	r.GET("/discussions", controller.FindDiscussions)
	r.GET("/discussion/:id", controller.FindDiscussion)
	r.POST("/discussion", controller.InsertDiscussion)
	r.PUT("/discussion/:id", controller.UpdateDiscussion)
	r.DELETE("/discussion/:id", controller.DeleteDiscussion)

	// comment context
	r.GET("/comments", controller.FindComments)
	r.GET("/comment/:id", controller.FindComment)
	r.POST("/comment", controller.InsertComment)
	r.PUT("/comment/:id", controller.UpdateComment)
	r.DELETE("/comment/:id", controller.DeleteComment)

	// uploader service
	r.POST("/upload", controller.UploadFile)

	// not found handler
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
		return
	})
	r.Run(models.GoDotEnvVariable("LISTEN_PORT"))
}
