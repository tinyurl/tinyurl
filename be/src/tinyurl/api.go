package main

import "github.com/gin-gonic/gin"
import "fmt"

func tinyUrlAPI(port string) {
	router := gin.Default()

	baseAPI := router.Group("/api/v1")
	{
		baseAPI.POST("/shorten", ShortenUrl)
		baseAPI.PUT("/health", HealthCheck)
	}

	router.GET("/", ParseUrl)
	router.GET("/:shortpath", ParseUrl)

	router.Run(port)
}

func ShortenUrl(c *gin.Context) {
	longurl := c.PostForm("longurl")

	// check longurl
	shortpath, exists := usi.dbs.CheckLongurl(longurl)
	if exists {
		c.JSON(200, gin.H{"status": 200, "shortpath": shortpath})
	} else {
		shortpath := usi.Shorten(longurl, 4)
		c.JSON(200, gin.H{"status": 200, "shorpath": shortpath})
	}
}

// ParseUrl parse shorten path and return source url
func ParseUrl(c *gin.Context) {
	shortpath := c.Param("shortpath")
	fmt.Println(shortpath)
	if len(shortpath) == 0 {
		c.Redirect(301, "https://adolphlwq.xyz")
	}

	longurl := usi.dbs.QueryUrlRecord(shortpath)
	if len(longurl) == 0 {
		c.Redirect(301, "https://adolphlwq.xyz")
	}

	c.Redirect(301, longurl)
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": 200, "data": "health"})
}
