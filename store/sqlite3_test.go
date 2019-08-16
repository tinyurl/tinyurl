package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configPathSqlite3                = "test_data/sqlite3.properties"
	sqlite3Client     *Sqlite3Client = NewSqlite3Client(configPathSqlite3)
)

func TestNewSqlite3Client(t *testing.T) {
	client := NewSqlite3Client(configPathSqlite3)
	assert.NotNil(t, client, "sqlite3 client should not be nil")
	// clean tmp file
	client.DropDatabase()
}
