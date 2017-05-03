package main

import(
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func tinyUrlAPI(port string) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://tinyurl.api.adolphlwq.xyz"},
		AllowMethods: []string{"*"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", 
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		//AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			fmt.Println(origin)
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	baseAPI := router.Group("/api/v1")
	{
		baseAPI.POST("/shorten", ShortenUrl)
		baseAPI.PUT("/health", HealthCheck)
	}

	//router.GET("/", ParseUrl)
	router.GET("/n/:shortpath", ParseUrl)

	router.Run(port)
}

func ShortenUrl(c *gin.Context) {
	longurl := c.PostForm("longurl")
	// check longurl
	logq.Info("check if longurl:", longurl, " has existed in db.")
	shortpath, exists := usi.dbs.CheckLongurl(longurl)
	if exists {
		logq.Info(longurl, " has been existed, return shortpath directly.")
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "shortpath": shortpath})
	} else {
		shortpath := usi.Shorten(longurl, 4)
		logq.Info("generate shortpath: ", shortpath, " for longurl: ", longurl)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "shorpath": shortpath})
	}
}

// ParseUrl parse shorten path and return source url
func ParseUrl(c *gin.Context) {
	shortpath := c.Param("shortpath")
	logq.Info("parse shortpath: ", shortpath, " for longurl")
	if len(shortpath) == 0 {
		logq.Warn("shortpath is nil, return default home path.")
		c.Redirect(http.StatusMovedPermanently, "http://tinyurl.adolphlwq.xyz")
	}

	longurl := usi.dbs.QueryUrlRecord(shortpath)
	if len(longurl) == 0 {
		logq.Warn("longurl of shortpath is nil, return default home page.")
		c.Redirect(http.StatusMovedPermanently, "http://tinyurl.adolphlwq.xyz")
	}

	c.Redirect(http.StatusMovedPermanently, longurl)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "health"})
}
