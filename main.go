package main

import (
	// "fmt"

	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

func GetSampleBlockchain(c *gin.Context) {
	resp, err := http.Get("https://api.blockchair.com/stats")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data interface{}
	json.Unmarshal(body, &data)

	c.JSON(http.StatusOK, data)
}

func main() {
	router := gin.Default()
	router.GET("/", GetSampleBlockchain)
	router.Run(":8080")
}
