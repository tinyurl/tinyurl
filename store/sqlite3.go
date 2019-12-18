package store

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/tinyurl/tinyurl/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Sqlite3Client struct {
	db         *gorm.DB
	DBName     string
	configPath string
}

func NewSqlite3Client(configPath string) *Sqlite3Client {
	setting := entity.GetGlobalConfig(configPath)
	source := fmt.Sprintf("%s/%s", setting.DBPath, setting.DBName)
	c := &Sqlite3Client{}
	db, err := gorm.Open("sqlite3", source)
	if err != nil {
		logrus.Fatalf("open connection to sqlite3 use gorm error: %s", err)
	}

	c.DBName = setting.DBName
	c.db = db
	c.configPath = configPath
	c.db.AutoMigrate(&entity.URL{})

	return c
}

func (c *Sqlite3Client) Create(url *entity.URL) {
	c.db.Create(url)
}

func (c *Sqlite3Client) Update(url *entity.URL) {
	c.db.Save(url)
}

func (c *Sqlite3Client) GetByOriginURL(originURL string) *entity.URL {
	var url entity.URL
	c.db.Where("origin_url = ?", originURL).First(&url)

	return &url
}

func (c *Sqlite3Client) GetByShortPath(shortPath string) *entity.URL {
	var url entity.URL
	c.db.Where("short_path = ?", shortPath).First(&url)

	return &url
}

// CreateDB check if database existed in db
// create database if not
func (c *Sqlite3Client) CreateDB(setting *entity.GlobalConfig) {
	source := fmt.Sprintf("%s/%s", setting.DBPath, setting.DBName)
	db, err := sql.Open("sqlite3", source)
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
func (c *Sqlite3Client) DropDatabase() {
	// sqlite does not have DROP DATABASE command, we just delete file
	setting := entity.GetGlobalConfig(c.configPath)
	source := fmt.Sprintf("%s/%s", setting.DBPath, setting.DBName)
	if err := os.Remove(source); err != nil {
		logrus.Fatalf("drop database %s error: %v", c.DBName, err)
	}
}
