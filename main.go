package main

import (
	"flag"
	"net"

	"github.com/tinyurl/tinyurl/entity"
	"github.com/tinyurl/tinyurl/server"
	"github.com/tinyurl/tinyurl/store"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.Parse()

	urlStore := store.NewGeneralDBClient(configPath)
	globalConfig := entity.GetGlobalConfigByViper(configPath)
	keyGenerater := entity.NewKeyGenerater(globalConfig.KeyAlgo)
	app := &entity.ServiceProvider{
		StoreClient:  urlStore,
		KeyGenerater: keyGenerater,
		GlobalConfig: globalConfig,
	}

	addr := net.JoinHostPort(app.GlobalConfig.Host, app.GlobalConfig.Port)
	server.Start(addr, app)
}
