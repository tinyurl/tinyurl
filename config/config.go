package config

import (
	"github.com/Sirupsen/logrus"
	"github.com/magiconair/properties"
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
