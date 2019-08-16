package server

import (
	"github.com/adolphlwq/tinyurl/config"
	"github.com/adolphlwq/tinyurl/store"
	"github.com/adolphlwq/tinyurl/uriuuid"
)

// ServiceProvider object hold service which server need
type ServiceProvider struct {
	MysqlClient  *store.Client
	UriUUID      uriuuid.UriUUID
	GlobalConfig *config.GlobalConfig
}
