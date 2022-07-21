package main

import (
	"fmt"

	"github.com/EKivutha/Url-Shortener/handler"
	"github.com/EKivutha/Url-Shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //initialize the socket
	//get func
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})
	r.GET("/:shorturl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()
	// fmt.Printf("Hello Go URL Shortener !🚀")
	// r.Run(":9808")
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
