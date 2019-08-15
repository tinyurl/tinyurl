package mysql

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/tinyurl/config"
	"github.com/adolphlwq/tinyurl/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Client hold mysql connection and wrape CRUD methods
type Client struct {
	DB       *gorm.DB
	Database string
}

// NewMySQLClient return new MySQLClient instance
func NewMySQLClient(configPath string) *Client {
	setting := config.GetGlobalConfig(configPath)
	mc := &Client{}

	CheckDB(setting)
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DBUser, setting.DBPassword, setting.DBHost, setting.DBPort, setting.DBName)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("open connection to mysql use gorm error: %s", err)
	}

	mc.Database = setting.DBName
	mc.DB = db
	mc.DB.AutoMigrate(&entity.URL{})
	logrus.Infof("create table urls done.\n")

	return mc
}

// CheckDB check if database existed in db
// create database if not
func CheckDB(setting *config.GlobalConfig) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", setting.DBUser, setting.DBPassword,
		setting.DBHost, setting.DBPort)
	db, err := sql.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("connection to db error: %s", err)
	}
	defer db.Close()

	sql := "CREATE DATABASE IF NOT EXISTS " + setting.DBName + ";"
	_, err = db.Exec(sql)
	if err != nil {
		logrus.Fatalf("create db %s error: %v", setting.DBName, err)
	}
}

// DropDatabase drop self hold database
func (c *Client) DropDatabase() {
	sql := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", c.Database)
	db := c.DB.DB()

	_, err := db.Exec(sql)
	if err != nil {
		logrus.Fatalf("drop database %s error: %v", c.Database, err)
	}
}
