package main

import (
	"flag"
	"net"

	"github.com/adolphlwq/tinyurl/entity"
	"github.com/adolphlwq/tinyurl/server"
	"github.com/adolphlwq/tinyurl/store"
	"github.com/adolphlwq/tinyurl/uriuuid"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.Parse()

	urlStore := store.GetURLStore(configPath)
	app := &entity.ServiceProvider{
		StoreClient:  urlStore,
		UriUUID:      uriuuid.BasicURIUUID{},
		GlobalConfig: entity.GetGlobalConfig(configPath),
	}

	addr := net.JoinHostPort(app.GlobalConfig.Host, app.GlobalConfig.Port)
	server.Start(addr, app)
}
