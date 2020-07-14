package main

import (
	"flag"
	"math"
	"net"

	_ "github.com/tinyurl/tinyurl/docs"
	"github.com/tinyurl/tinyurl/domain"
	"github.com/tinyurl/tinyurl/server"
	"github.com/tinyurl/tinyurl/store"
)

// @title TinyURL Swagger API Docs
// @version 1.0
// @description TinyURL API Document

// @contact.name AdolphLWQ
// @contact.url https://git.io/tinyurl
// @contact.email kenan3015@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.Parse()

	generalStore := store.NewGeneralDBClient(configPath)
	globalConfig := domain.GetGlobalConfigByViper(configPath)
	keyGenerater := domain.NewKeyGenerater(globalConfig.KeyAlgo)

	switch globalConfig.KeyAlgo {
	case domain.KeyAlgoSender:
		sender := generalStore.GetSenderWorker()
		if sender.Index != 0 {
			keyGenerater.SetIndex(sender.Index)
		} else {
			// init start
			var index int64
			index = int64(math.Pow(domain.DefaultCharsLen, float64(globalConfig.KeyLen-1)))
			keyGenerater.SetIndex(index)
			generalStore.UpdateSenderWorker(sender)
		}
	}
	app := &domain.ServiceProvider{
		StoreClient:  generalStore,
		KeyGenerater: keyGenerater,
		GlobalConfig: globalConfig,
	}

	addr := net.JoinHostPort(app.GlobalConfig.Host, app.GlobalConfig.Port)
	server.Start(addr, app)
}
