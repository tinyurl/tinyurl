package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProps(t *testing.T) {
	configPath := "../defult.properties"
	props := ReadProps(configPath)
	assert.NotNil(t, props)

	database := props.MustGetString("db.name")
	assert.Equal(t, database, "tinyurldb")
}

func TestGetGlobalConfig(t *testing.T) {
	configPath := "../defult.properties"
	props := ReadProps(configPath)
	assert.NotNil(t, props)

	cc := GetGlobalConfig(configPath)
	assert.Equal(t, cc.Host, "0.0.0.0", "host should be 0.0.0.0")
	assert.Equal(t, cc.Port, "8877", "port should be 8877")
	assert.Equal(t, cc.DBType, "sqlite3", "db type should be sqlite3")
	assert.Equal(t, cc.DBHost, "127.0.0.1", "db host should be 127.0.0.1")
	assert.Equal(t, cc.DBPort, "3306", "db port should be 3306")
	assert.Equal(t, cc.DBName, "tinyurldb", "db name should be tinyurldb")
}
