package main

import (
	"flag"
	"net"

	"github.com/tinyurl/tinyurl/entity"
	"github.com/tinyurl/tinyurl/server"
	"github.com/tinyurl/tinyurl/store"
	"github.com/tinyurl/tinyurl/uriuuid"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.Parse()

	urlStore := store.GetURLStore(configPath)
	app := &entity.ServiceProvider{
		StoreClient:  urlStore,
		UriUUID:      uriuuid.BasicURIUUID{},
		GlobalConfig: entity.GetGlobalConfigByViper(configPath),
	}

	addr := net.JoinHostPort(app.GlobalConfig.Host, app.GlobalConfig.Port)
	server.Start(addr, app)
}
