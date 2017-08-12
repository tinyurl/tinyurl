package server

import "github.com/adolphlwq/tinyurl/mysql"

// ServiceProvider object hold service which server need
type ServiceProvider struct {
	MysqlClient *mysql.Client
}
