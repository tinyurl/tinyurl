package main

import(
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status":200, "data": "pong"})
	})

	router.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": 200, "data":"health"})
	})

	router.Run(":8888")
}
