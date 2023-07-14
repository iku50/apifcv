package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iku50/apifcv/src/api"
)

type CV struct {
	Content string `json:"content" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/gpt", handler)
	router.POST("/cv", CVHandler)
	api.Init()

	if err := router.Run(":12321"); err != nil {
		log.Fatal(err.Error())
	}
}

func handler(c *gin.Context) {
	role := c.Query("role")
	prompt := c.Query("prompt")
	fmt.Println(role, prompt)
	c.JSON(200, gin.H{
		"message": api.GptGet(role, prompt),
	})
}

func CVHandler(c *gin.Context) {
	var CV CV
	if err := c.ShouldBindJSON(&CV); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(CV)
	cvtext := CV.Content
	c.JSON(200, gin.H{
		"message": api.CVGet(cvtext),
	})
}
