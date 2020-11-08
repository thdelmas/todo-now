package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	fmt.Println("yolo")
	Router = gin.Default()
	Router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	Router.Run()
}
