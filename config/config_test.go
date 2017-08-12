package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProps(t *testing.T) {
	configPath := "../test.properties"
	props := ReadProps(configPath)
	assert.NotNil(t, props)

	database := props.MustGetString("mysql.database")
	assert.Equal(t, database, "test_tinyurl")
}
