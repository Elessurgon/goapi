package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/", pong)
	router.POST("/send-json", search_data)
	
	router.Run()
}