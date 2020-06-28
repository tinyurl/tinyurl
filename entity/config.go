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
	Host                string
	Port                string
	Domain              string
	SwaggerURL          string
	KeyAlgo             string
	KeyBasicDefaultLen  int
	KeySenderDefaultLen int

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
	host := props.MustGet("app.host")
	port := props.MustGet("app.port")
	swaggerURL := fmt.Sprintf("%s:%s/swagger/doc.json", host, port)

	config := &GlobalConfig{
		Host:       host,
		Port:       port,
		Domain:     props.MustGet("app.domain"),
		SwaggerURL: swaggerURL,
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

	host := viper.GetString("app.host")
	port := viper.GetString("app.port")
	swaggerURL := fmt.Sprintf("%s:%s/swagger/doc.json", host, port)

	config := &GlobalConfig{
		Host:                host,
		Port:                port,
		Domain:              viper.GetString("app.domain"),
		SwaggerURL:          swaggerURL,
		KeyAlgo:             viper.GetString("app.key.algorithm"),
		KeyBasicDefaultLen:  viper.GetInt("app.basic.default.len"),
		KeySenderDefaultLen: viper.GetInt("app.sender.default.len"),
		DBType:              viper.GetString("db.type"),
		DBPath:              viper.GetString("db.path"),
		DBHost:              viper.GetString("db.host"),
		DBPort:              viper.GetString("db.port"),
		DBName:              viper.GetString("db.name"),
		DBUser:              viper.GetString("db.user"),
		DBPassword:          viper.GetString("db.password"),
	}

	// validate key generater algorithm and config
	switch config.KeyAlgo {
	case KeyAlgoRandom:
		if config.KeyBasicDefaultLen <= 0 {
			panic(fmt.Errorf("key algo is %s, default len is %d, should >= 0",
				config.KeyAlgo, config.KeyBasicDefaultLen))
		}
	case KeyAlgoSender:
		if config.KeySenderDefaultLen <= 0 {
			panic(fmt.Errorf("key algo is %s, default len is %d, should >= 0",
				config.KeyAlgo, config.KeySenderDefaultLen))
		}
	}

	return config
}

// db type
const (
	SQLITE3 = "sqlite3"
	MYSQL   = "mysql"
)
