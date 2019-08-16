package store

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/tinyurl/config"
)

var (
	configPath                      = "../defult.properties"
	setting    *config.GlobalConfig = config.GetGlobalConfig(configPath)
	client     *MySQLClient         = NewMySQLClient(configPath)
)

func newSqlDB(setting *config.GlobalConfig) *sql.DB {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", setting.DBUser, setting.DBPassword,
		setting.DBHost, setting.DBPort)
	db, err := sql.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("connection to db error: %s", err)
	}

	return db
}

func TestNewMySQLClient(t *testing.T) {
	client := NewMySQLClient(configPath)
	if client == nil {
		t.Errorf("client should not be nil")
	}

	t.Logf("new mysql client success, drop test database.\n")
	client.DropDatabase()
}

func TestCreateDB(t *testing.T) {
	client.CreateDB(setting)
	db := newSqlDB(setting)
	defer db.Close()

	// check if database exist
	sql := fmt.Sprintf("USE %s;", setting.DBName)
	_, err := db.Exec(sql)
	if err != nil {
		t.Errorf("show databases error: %s\n", err)
	}

	t.Logf("init db success, drop test database.\n")
	client.DropDatabase()
}

func TestDropDatabase(t *testing.T) {
	client.DropDatabase()
}
