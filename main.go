package main

import (
	"flag"
	"net"

	"github.com/adolphlwq/tinyurl/config"
	"github.com/adolphlwq/tinyurl/server"
	"github.com/adolphlwq/tinyurl/store"
	"github.com/adolphlwq/tinyurl/uriuuid"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.Parse()

	mysqlClient := store.NewMySQLClient(configPath)
	app := &server.ServiceProvider{
		MysqlClient:  mysqlClient,
		UriUUID:      uriuuid.BasicURIUUID{},
		GlobalConfig: config.GetGlobalConfig(configPath),
	}

	addr := net.JoinHostPort(app.GlobalConfig.Host, app.GlobalConfig.Port)
	server.Start(addr, app)
}
