package server

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/tinyurl/entity"
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
	if url.OriginURL == "" {
		url.CreateTime = time.Now().UTC()
		url.OriginURL = OriginURL
		url.ShortPath = "not implement"
		c.JSON(http.StatusOK, gin.H{
			"short_path": url.ShortPath,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"short_path": url.ShortPath,
		})
	}
}
