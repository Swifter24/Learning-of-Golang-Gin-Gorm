package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	r.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "put")
	})
	r.DELETE("/del", func(c *gin.Context) {
		c.String(http.StatusOK, "delete")
	})
	err := r.Run()
	if err != nil {
		return
	}
}
