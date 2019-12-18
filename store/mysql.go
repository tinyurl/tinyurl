package store

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/tinyurl/tinyurl/entity"
)

type MySQLClient struct {
	db     *gorm.DB
	DBName string
}

func NewMySQLClient(configPath string) *MySQLClient {
	setting := entity.GetGlobalConfig(configPath)
	c := &MySQLClient{}
	c.CreateDB(setting)

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DBUser, setting.DBPassword, setting.DBHost, setting.DBPort, setting.DBName)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("open connection to mysql use gorm error: %s", err)
	}

	c.DBName = setting.DBName
	c.db = db
	c.db.AutoMigrate(&entity.URL{})

	return c
}

func (c *MySQLClient) Create(url *entity.URL) {
	c.db.Create(url)
}

func (c *MySQLClient) Update(url *entity.URL) {
	c.db.Save(url)
}

func (c *MySQLClient) GetByOriginURL(originURL string) *entity.URL {
	var url entity.URL
	c.db.Where("origin_url = ?", originURL).First(&url)

	return &url
}

func (c *MySQLClient) GetByShortPath(shortPath string) *entity.URL {
	var url entity.URL
	c.db.Where("short_path = ?", shortPath).First(&url)

	return &url
}

// CreateDB check if database existed in db
// create database if not
func (c *MySQLClient) CreateDB(setting *entity.GlobalConfig) {
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
func (c *MySQLClient) DropDatabase() {
	sql := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", c.DBName)
	db := c.db.DB()

	_, err := db.Exec(sql)
	if err != nil {
		logrus.Fatalf("drop database %s error: %v", c.DBName, err)
	}
}
