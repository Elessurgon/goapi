package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/ping", pong)
	router.POST("/send-json", search_data)
	
	router.Run()
}