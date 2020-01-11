package entity

import (
	"fmt"
	"path"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/magiconair/properties"
	"github.com/spf13/viper"
)

// ReadProps read and return props according configPath
func ReadProps(configPath string) *properties.Properties {
	props, err := properties.LoadFile(configPath, properties.UTF8)
	if err != nil {
		logrus.Fatalf("read %s error: %v\n", configPath, err)
	}

	return props
}

type GlobalConfig struct {
	// app config
	Host   string
	Port   string
	Domain string

	// DB config
	DBType     string
	DBPath     string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

// GetGlobalConfig use properties
func GetGlobalConfig(configPath string) *GlobalConfig {
	props := ReadProps(configPath)

	config := &GlobalConfig{
		Host:       props.MustGet("app.host"),
		Port:       props.MustGet("app.port"),
		Domain:     props.MustGet("app.domain"),
		DBType:     props.MustGet("db.type"),
		DBPath:     props.GetString("db.path", ""),
		DBHost:     props.GetString("db.host", ""),
		DBPort:     props.GetString("db.port", ""),
		DBName:     props.GetString("db.name", ""),
		DBUser:     props.GetString("db.user", ""),
		DBPassword: props.GetString("db.password", ""),
	}

	return config
}

// GetGlobalConfigByViper use viper
func GetGlobalConfigByViper(configPath string) *GlobalConfig {
	baseConfigPath := path.Base(configPath)
	configName := strings.Split(baseConfigPath, ".")[0]

	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := &GlobalConfig{
		Host:       viper.GetString("app.host"),
		Port:       viper.GetString("app.port"),
		Domain:     viper.GetString("app.domain"),
		DBType:     viper.GetString("db.type"),
		DBPath:     viper.GetString("db.path"),
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBName:     viper.GetString("db.name"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: viper.GetString("db.password"),
	}

	return config
}

// db type
const (
	SQLITE3 = "sqlite3"
	MYSQL   = "mysql"
)
