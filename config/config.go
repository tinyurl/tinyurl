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
