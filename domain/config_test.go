package domain

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestReadProps(t *testing.T) {
	configPath := "../default.properties"
	props := ReadProps(configPath)
	assert.NotNil(t, props)

	database := props.MustGetString("db.name")
	assert.Equal(t, database, "tinyurldb")
}

func TestGetGlobalConfig(t *testing.T) {
	configPathDefault := "test_data/default.properties"
	configPathMySQL := "test_data/mysql.properties"
	propsDefault := ReadProps(configPathDefault)
	propsMySQL := ReadProps(configPathMySQL)
	c := GetGlobalConfig(configPathDefault)
	cm := GetGlobalConfig(configPathMySQL)

	assert.NotNil(t, propsDefault)
	assert.NotNil(t, propsMySQL)

	assert.Equal(t, c.Host, "0.0.0.0", "host should be 0.0.0.0")
	assert.Equal(t, c.Port, "8877", "port should be 8877")
	assert.Equal(t, c.DBType, "sqlite3", "db type should be sqlite3")
	assert.Equal(t, c.DBHost, "", "db host should be empty")
	assert.Equal(t, c.DBPath, ".", "db path should be .")
	assert.Equal(t, c.DBName, "tinyurldb", "db name should be tinyurldb")

	assert.Equal(t, cm.Host, "0.0.0.0", "host should be 0.0.0.0")
	assert.Equal(t, cm.Port, "8877", "port should be 8877")
	assert.Equal(t, cm.DBType, "mysql", "db type should be mysql")
	assert.Equal(t, cm.DBHost, "127.0.0.1", "db host should be 127.0.0.1")
	assert.Equal(t, cm.DBPort, "3306", "db port should be 3306")
	assert.Equal(t, cm.DBName, "tinyurldb", "db name should be tinyurldb")
	assert.Equal(t, cm.DBUser, "root", "db name should be root")
}

func TestGetGlobalConfigByViper(t *testing.T) {
	configPathDefault := "test_data/default.properties"
	configPathMySQL := "test_data/mysql.properties"
	viper.AddConfigPath("test_data")
	c := GetGlobalConfigByViper(configPathDefault)
	cm := GetGlobalConfigByViper(configPathMySQL)

	assert.Equal(t, c.Host, "0.0.0.0", "host should be 0.0.0.0")
	assert.Equal(t, c.Port, "8877", "port should be 8877")
	assert.Equal(t, c.KeyAlgo, "random", "KeyAlgo should be random")
	assert.Equal(t, c.KeyLen, 6, "KeyLen should be 6")
	assert.Equal(t, c.DBType, "sqlite3", "db type should be sqlite3")
	assert.Equal(t, c.DBHost, "", "db host should be empty")
	assert.Equal(t, c.DBPath, ".", "db path should be .")
	assert.Equal(t, c.DBName, "tinyurldb", "db name should be tinyurldb")

	assert.Equal(t, cm.Host, "0.0.0.0", "host should be 0.0.0.0")
	assert.Equal(t, cm.Port, "8877", "port should be 8877")
	assert.Equal(t, cm.DBType, "mysql", "db type should be mysql")
	assert.Equal(t, cm.DBHost, "127.0.0.1", "db host should be 127.0.0.1")
	assert.Equal(t, cm.DBPort, "3306", "db port should be 3306")
	assert.Equal(t, cm.DBName, "tinyurldb", "db name should be tinyurldb")
	assert.Equal(t, cm.DBUser, "root", "db name should be root")
}
