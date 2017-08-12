package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
)

var (
	setting *MySQLSetting = NewMySQLSetting(configPath)
	client  *Client       = NewMySQLClient(configPath)
)

func newSqlDB(setting *MySQLSetting) *sql.DB {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", setting.User, setting.Password,
		setting.Host, setting.Port)
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

func TestCheckDB(t *testing.T) {
	CheckDB(setting)
	db := newSqlDB(setting)
	defer db.Close()

	// check if database exist
	sql := fmt.Sprintf("USE %s;", setting.Database)
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
