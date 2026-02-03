package main

import "github.com/gin-gonic/gin"

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildTime = "unknown"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version":   Version,
			"gitCommit": GitCommit,
			"buildTime": BuildTime,
		})
	})
	router.Run()
}
