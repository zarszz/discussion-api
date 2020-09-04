package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", controller.FindUsers)
	r.Run()
}
