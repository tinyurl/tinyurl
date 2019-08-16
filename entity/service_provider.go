package entity

import (
	"github.com/adolphlwq/tinyurl/config"
	"github.com/adolphlwq/tinyurl/uriuuid"
)

type URLStore interface {
	Create(url *URL)
	Update(url *URL)
	GetByOriginURL(originURL string) *URL
	GetByShortPath(shortPath string) *URL
	DropDatabase()
}

// ServiceProvider object hold service which server need
type ServiceProvider struct {
	StoreClient  URLStore
	UriUUID      uriuuid.UriUUID
	GlobalConfig *config.GlobalConfig
}
