package server

import (
	"testing"
	"time"

	"github.com/adolphlwq/tinyurl/mysql"
)

const (
	TestPort   = "9090"
	TestAddr   = "http://0.0.0.0:9090"
	ConfigPath = "../test.properties"
)

var (
	mysqlClient *mysql.Client
	appService  *ServiceProvider
)

func init() {
	mysqlClient = mysql.NewMySQLClient(ConfigPath)
	appService = &ServiceProvider{
		MysqlClient: mysqlClient,
	}
}

// startTestServer
func startTestServer(t *testing.T) {
	go func() {
		r := BuildEngine(appService)
		r.Run(":" + TestPort)
	}()

	t.Logf("wait 2s to start testServer...\n")
	time.Sleep(time.Second * 2)
}
