package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinyurl/tinyurl/entity"
)

var (
	configPathMySQL   = "test_data/mysql.properties"
	configPathSqlite3 = "test_data/sqlite3.properties"
)

func TestInitDB(t *testing.T) {
	t.Log("test mysql InitDB...")
	settingMySQL := entity.GetGlobalConfig(configPathMySQL)
	InitDB(settingMySQL)
	NewGeneralDBClient(configPathMySQL).DropDatabase()
	t.Logf("test mysql InitDB success.\n")

	t.Log("test sqlite3 InitDB...")
	settingSqlite3 := entity.GetGlobalConfig(configPathSqlite3)
	InitDB(settingSqlite3)
	NewGeneralDBClient(configPathSqlite3).DropDatabase()
	t.Logf("test sqlite3 InitDB success.\n")
}

func TestNewMySQLClient(t *testing.T) {
	client := NewGeneralDBClient(configPathMySQL)
	if client == nil {
		t.Errorf("client should not be nil")
	}

	t.Logf("new mysql client success, drop test database.\n")
	client.DropDatabase()
}

func TestNewSqlite3Client(t *testing.T) {
	client := NewGeneralDBClient(configPathSqlite3)
	assert.NotNil(t, client, "sqlite3 client should not be nil")
	// clean tmp file
	client.DropDatabase()
}

func TestDropDatabase(t *testing.T) {
	NewGeneralDBClient(configPathMySQL).DropDatabase()
	NewGeneralDBClient(configPathSqlite3).DropDatabase()
}

func TestUpdateSenderWorker(t *testing.T) {
	// this process will create init sender with index=0
	client := NewGeneralDBClient(configPathSqlite3)
	sender := client.GetSenderWorker()

	sender.Index = 7
	client.UpdateSenderWorker(sender)
	ret := client.GetSenderWorker()

	assert.Equal(t, sender.Index, ret.Index, "index should be %d, but get %d", sender.Index, ret.Index)
	client.DropDatabase()
}
