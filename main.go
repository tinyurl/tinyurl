package main

import (
	"flag"
	"net"

	"github.com/adolphlwq/tinyurl/mysql"
	"github.com/adolphlwq/tinyurl/server"
)

func main() {
	var (
		configPath string
		host       string
		port       string
	)
	flag.StringVar(&configPath, "config", "default.properties", "config path")
	flag.StringVar(&host, "host", "0.0.0.0", "tinyurl server bind host")
	flag.StringVar(&port, "port", "8877", "tinyurl server bind port")
	flag.Parse()

	mysqlClient := mysql.NewMySQLClient(configPath)
	sp := &server.ServiceProvider{
		MysqlClient: mysqlClient,
	}
	addr := net.JoinHostPort(host, port)

	server.Start(addr, sp)
}
