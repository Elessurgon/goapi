package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func search_data(c *gin.Context) {
	var jsonData []map[string]interface{}
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
        return
	}
	// fmt.Println(string(data))
	if err := json.Unmarshal(data, &jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
        return
	}

	transformedData := searchData(jsonData, "SearchTerms", ",")
	c.JSON(http.StatusOK, transformedData)
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping-pong",
	})
}