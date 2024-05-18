package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhil0020/gin/02_middleware/middleware"
)

func main() {
	router := gin.Default()

	router.GET("/getData", middleware.Authenticate, getData)
	router.GET("/getData1", getData1)
	router.GET("/getData2", getData2)

	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GetData Method",
	})
}

func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GetData 1 Method",
	})
}

func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GetData 2 Method",
	})
}
