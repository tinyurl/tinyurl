package server

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/tinyurl/entity"
	"github.com/gin-gonic/gin"
)

// ShortenURL shorten origin url and save to db
func ShortenURL(c *gin.Context, sp *ServiceProvider) {
	originUrl := c.PostForm("origin_url")

	if originUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please provide origin_url.",
		})
		return
	}
	// check longurl
	logrus.Infof("check if origin %s has existed in db.\n", originUrl)
	var url entity.URL
	sp.MysqlClient.DB.Where("origin_url = ?", originUrl).First(&url)
	if url.OriginUrl == "" {
		url.CreateTime = time.Now().UTC()
		url.OriginUrl = originUrl
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
