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
		DBHost:     props.MustGetString("db.host"),
		DBPort:     props.MustGetString("db.port"),
		DBName:     props.MustGetString("db.name"),
		DBUser:     props.MustGetString("db.user"),
		DBPassword: props.MustGetString("db.password"),
	}

	return config
}
