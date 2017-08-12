package mysql

import (
	"testing"

	"github.com/adolphlwq/tinyurl/config"
	"github.com/stretchr/testify/assert"
)

var configPath = "../test.properties"

func TestNewMySQLSetting(t *testing.T) {
	mysqlSetting := NewMySQLSetting(configPath)
	assert.NotNil(t, mysqlSetting)
	assert.Equal(t, mysqlSetting.Database, "test_tinyurl")
}

func TestNewMySQLSettingFromConfig(t *testing.T) {
	props := config.ReadProps(configPath)
	mysqlSetting := NewMySQLSettingFromConfig(props)
	assert.NotNil(t, mysqlSetting)
	assert.Equal(t, mysqlSetting.Database, "test_tinyurl")
}
