package store

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/tinyurl/tinyurl/entity"
)

var (
	configPathMySQL              = "test_data/mysql.properties"
	mysqlClient     *MySQLClient = NewMySQLClient(configPathMySQL)
)

func newSqlDB(configPath string) *sql.DB {
	setting := entity.GetGlobalConfig(configPath)
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", setting.DBUser, setting.DBPassword,
		setting.DBHost, setting.DBPort)
	db, err := sql.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("connection to db error: %s", err)
	}

	return db
}

func TestNewMySQLClient(t *testing.T) {
	client := NewMySQLClient(configPathMySQL)
	if client == nil {
		t.Errorf("client should not be nil")
	}

	t.Logf("new mysql client success, drop test database.\n")
	client.DropDatabase()
}

func TestCreateDB(t *testing.T) {
	setting := entity.GetGlobalConfig(configPathMySQL)
	mysqlClient.CreateDB(setting)
	db := newSqlDB(configPathMySQL)
	defer db.Close()

	// check if database exist
	sql := fmt.Sprintf("USE %s;", setting.DBName)
	_, err := db.Exec(sql)
	if err != nil {
		t.Errorf("show databases error: %s\n", err)
	}

	t.Logf("init db success, drop test database.\n")
	mysqlClient.DropDatabase()
}

func TestDropDatabase(t *testing.T) {
	mysqlClient.DropDatabase()
}
