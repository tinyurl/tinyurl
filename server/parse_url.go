package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/tinyurl/entity"
	"github.com/gin-gonic/gin"
)

// ParseURL parse shorten path and return source url
func ParseURL(c *gin.Context, appService *ServiceProvider) {
	shortPath := c.Param("shortpath")
	if shortPath == "" {
		logrus.Warnf("shortpath is nil, return default home path.\n")
		c.Redirect(http.StatusMovedPermanently, appService.GlobalConfig.Domain)
	}

	var url entity.URL
	appService.MysqlClient.DB.Where("short_path = ?", shortPath).First(&url)
	if url.OriginURL == "" {
		logrus.Warnf("short url has no record in db.\n")
		c.Redirect(http.StatusMovedPermanently, appService.GlobalConfig.Domain)
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginURL)
}
