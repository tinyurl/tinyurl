package store

import (
	"github.com/adolphlwq/tinyurl/entity"
	_ "github.com/go-sql-driver/mysql"
)

func GetURLStore(configPath string) entity.URLStore {
	var urlStore entity.URLStore
	setting := entity.GetGlobalConfig(configPath)
	switch setting.DBType {
	case "mysql":
		urlStore = NewMySQLClient(configPath)
		break
	case "sqlite3":
		urlStore = NewSqlite3Client(configPath)
		break
	default:
		urlStore = NewMySQLClient(configPath)
	}

	return urlStore
}
