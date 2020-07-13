package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tinyurl/tinyurl/domain"
)

// ParseURL parse shorten path and return source url
func ParseURL(c *gin.Context, appService *domain.ServiceProvider) {
	shortPath := c.Param("shortpath")
	if shortPath == "" {
		logrus.Warnf("shortpath is nil, return default home path.\n")
		c.Redirect(http.StatusMovedPermanently, appService.GlobalConfig.Domain)
	}

	url := appService.StoreClient.GetByShortPath(shortPath)
	if url.OriginURL == "" {
		logrus.Warnf("short url has no record in db.\n")
		c.Redirect(http.StatusMovedPermanently, appService.GlobalConfig.Domain)
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginURL)
}
