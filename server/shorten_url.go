package server

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/tinyurl/entity"
	"github.com/adolphlwq/tinyurl/mysql"
	"github.com/gin-gonic/gin"
)

// ShortenURL shorten origin url and save to db
func ShortenURL(c *gin.Context, appService *ServiceProvider) {
	OriginURL := c.PostForm("origin_url")

	if OriginURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please provide origin_url.",
		})
		return
	}

	// check longurl
	logrus.Infof("check if origin %s has existed in db.\n", OriginURL)
	var url entity.URL
	appService.MysqlClient.DB.Where("origin_url = ?", OriginURL).First(&url)
	// url.OriginURL == "" OriginURL does not exist in db, generate shortpath
	if url.OriginURL == "" {
		var retry = 0
		shortPath := appService.UriUUID.New()
		for isShortPathExisted(appService.MysqlClient, shortPath) {
			shortPath = appService.UriUUID.New()
			retry++
		}
		logrus.Infof("generate short path for %s with %d times\n", OriginURL, retry)

		url.CreateTime = time.Now().UTC()
		url.OriginURL = OriginURL
		url.ShortPath = shortPath
		appService.MysqlClient.DB.Create(&url)

		c.JSON(http.StatusOK, gin.H{
			"message":    ShortenURLSuccess,
			"short_path": url.ShortPath,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    ShortPathExisted,
			"short_path": url.ShortPath,
		})
	}
}

// isShortPathExisted true short path exists or false not exist
func isShortPathExisted(client *mysql.Client, shortPath string) bool {
	var url entity.URL
	client.DB.Where("short_path = ?", shortPath).First(&url)
	return url.OriginURL != ""
}
