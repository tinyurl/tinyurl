package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/tinyurl/tinyurl/entity"
)

// ShortenURL shorten origin url and save to db
func ShortenURL(c *gin.Context, appService *entity.ServiceProvider) {
	OriginURL := c.PostForm("origin_url")

	if OriginURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please provide origin_url.",
		})
		return
	}

	// check longurl
	logrus.Infof("check if origin url %s has existed in db.\n", OriginURL)
	url := appService.StoreClient.GetByOriginURL(OriginURL)
	if url.OriginURL == "" {
		url.ShortPath = appService.KeyGenerater.New()
		url.CreateTime = time.Now().UTC()
		url.OriginURL = OriginURL
		appService.StoreClient.Create(url)

		c.JSON(http.StatusOK, gin.H{
			"message":    ShortenURLSuccess,
			"short_path": fmt.Sprintf("%s/n/%s", appService.GlobalConfig.Domain, url.ShortPath),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    ShortPathExisted,
			"short_path": fmt.Sprintf("%s/n/%s", appService.GlobalConfig.Domain, url.ShortPath),
		})
	}
}
