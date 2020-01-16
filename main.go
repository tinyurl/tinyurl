package main

import (
	"flag"
	"math"
	"net"

	"github.com/tinyurl/tinyurl/entity"
	"github.com/tinyurl/tinyurl/server"
	"github.com/tinyurl/tinyurl/store"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.Parse()

	generalStore := store.NewGeneralDBClient(configPath)
	globalConfig := entity.GetGlobalConfigByViper(configPath)
	keyGenerater := entity.NewKeyGenerater(globalConfig.KeyAlgo)

	switch globalConfig.KeyAlgo {
	case entity.KeyAlgoSender:
		sender := generalStore.GetSenderWorker()
		if sender.Index != 0 {
			keyGenerater.SetIndex(sender.Index)
		} else {
			// init start
			var index int64
			index = int64(math.Pow(entity.DefaultCharsLen, float64(globalConfig.KeySenderDefaultLen-1)))
			keyGenerater.SetIndex(index)
			generalStore.UpdateSenderWorker(sender)
		}
	}
	app := &entity.ServiceProvider{
		StoreClient:  generalStore,
		KeyGenerater: keyGenerater,
		GlobalConfig: globalConfig,
	}

	addr := net.JoinHostPort(app.GlobalConfig.Host, app.GlobalConfig.Port)
	server.Start(addr, app)
}
