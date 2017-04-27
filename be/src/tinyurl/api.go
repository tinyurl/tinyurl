package main

import(
	"github.com/gin-gonic/gin"
)


func tinyUrlAPI(port string) {
	router := gin.Default()

	router.POST("/shorten", ShortenUrl)
	router.GET("/", ParseUrl)
	router.GET("/health", HealthCheck)

	router.Run(port)
}

func ShortenUrl(c *gin.Context) {
	longurl := c.PostForm("longurl")
	
	// check longurl
	shortpath, exists := usi.dbs.CheckLongurl(longurl)
	if exists {
		c.JSON(200, gin.H{"status": 200, "shortpath": shortpath})
	}else{
		shortpath := usi.Shorten(longurl, 4)
		c.JSON(200, gin.H{"status": 200, "shorpath": shortpath})
	}
}

func ParseUrl(c *gin.Context) {
	c.JSON(200, gin.H{"status": 200, "data": "parseurl"})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": 200, "data": "health"})
}
