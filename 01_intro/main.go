package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/getQueryString", getQueryString)

	// basic auth

	auth := gin.BasicAuth(gin.Accounts{
		"user1":  "pass1",
		"user2":  "pass2",
		"nikhil": "12345",
	})

	// here key is user name and value is password

	// Router Grouping

	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}

	client := router.Group("/client")
	{
		client.GET("/getDataPost", getDataPost)

	}

	// Custom server configs

	// router.Run(":3000")

	server := &http.Server{
		Addr:         ":5000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am a GIN framework",
	})
}

func getDataPost(c *gin.Context) {
	body := c.Request.Body

	values, _ := io.ReadAll(body)

	c.JSON(200, gin.H{
		"data": "Hi I am a GIN framework",
		"body": string(values),
	})
}

func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
