package mysql

import (
	"github.com/adolphlwq/tinyurl/config"
	"github.com/magiconair/properties"
)

// MySQLSetting wrape mysql setting
type MySQLSetting struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

// NewMySQLSettingFromConfig return new mysql setting from properties.Properties
func NewMySQLSettingFromConfig(props *properties.Properties) *MySQLSetting {
	ms := &MySQLSetting{
		Host:     props.MustGetString("mysql.host"),
		Port:     props.MustGetString("mysql.port"),
		Database: props.MustGetString("mysql.database"),
		User:     props.MustGetString("mysql.user"),
		Password: props.MustGetString("mysql.password"),
	}
	return ms
}

// NewMySQLSetting return new mysql setting from config file path
func NewMySQLSetting(configPath string) *MySQLSetting {
	props := config.ReadProps(configPath)
	return NewMySQLSettingFromConfig(props)
}
