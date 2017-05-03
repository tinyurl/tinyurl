package main

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

var onceSetupDB sync.Once

type Url struct {
	Id          int
	Longurl     string
	Shortpath   string
	CreatedTime time.Time
}

type DBService struct {
	DBName  string
	User    string
	Pass    string
	Address string
	Port    string
	DB      *sql.DB
}

func NewDB(dbname, user, pass, address, port string) *DBService {
	dbpath := user + ":" + pass + "@tcp(" + address + ":" + port + ")/"
	fmt.Println(dbpath)
	db, err := sql.Open("mysql", dbpath)
	if err != nil {
		logq.Fatal("setup up db error:", err)
	}

	dbs := &DBService{
		DBName:  dbname,
		User:    user,
		Pass:    pass,
		Address: address,
		Port:    port,
		DB:      db,
	}

	logq.Info("Start setup db ", dbs.DBName)
	dbs.Setup()
	return dbs
}

func (dbs *DBService) Close() {
	dbs.DB.Close()
}

func (dbs *DBService) Setup() {
	urlTable := `
		CREATE TABLE IF NOT EXISTS ` + dbs.DBName + `.url (
			id INT(10) NOT NULL AUTO_INCREMENT,
			longurl VARCHAR(21800) NOT NULL,
			shortpath VARCHAR(32) NOT NULL,
			created_time DATE NOT NULL,
			PRIMARY KEY (id)
		);
	`
	dbSchema := `
		CREATE DATABASE IF NOT EXISTS ` + dbs.DBName + ` 
			DEFAULT CHARACTER SET utf8
			DEFAULT COLLATE utf8_general_ci;
	`
	useDB := "USE " +dbs.DBName + ";"

	onceSetupDB.Do(func() {
		if _, err := dbs.DB.Exec(dbSchema); err != nil {
			logq.Fatal("setup database ", dbs.DBName, " err:", err)
		}
		
		if _, err := dbs.DB.Exec(useDB); err != nil {
			logq.Fatal("use db ", dbs.DBName, " error:", err)
		}

		if _, err := dbs.DB.Exec(urlTable); err != nil {
			logq.Fatal("setup table error:", err)
		}
	})
}

// CheckLongurl check if longurl has existed
// return false means longurl not exists in db
func (dbs *DBService) CheckLongurl(longurl string) (string, bool) {
	stmt, err := dbs.DB.Prepare("SELECT longurl, shortpath FROM url WHERE longurl=?")
	defer stmt.Close()
	if err != nil {
		logq.Fatal("prepare longurl stmt error: ", err)
	}

	var longurl_, shortpath string
	err = stmt.QueryRow(longurl).Scan(&longurl_, &shortpath)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false
		} else {
			logq.Fatal("query longurl ", longurl, " error: ", err)
		}
	}

	if len(shortpath) != 0 {
		return shortpath, true
	} else {
		return "", false
	}
}

// CheckPath check if shortpath has existed
// return false means shortpath not extsts in db
func (dbs *DBService) CheckPath(shortpath string) bool {
	stmt, err := dbs.DB.Prepare("SELECT shortpath FROM url WHERE shortpath=?")
	defer stmt.Close()
	if err != nil {
		logq.Fatal("check shortpath ", shortpath, " err:", err)
	}

	var ret string
	err = stmt.QueryRow(shortpath).Scan(&ret)
	if err != nil {
		//refer http://go-database-sql.org/errors.html
		if err == sql.ErrNoRows {
			return false
		} else {
			logq.Fatal("check if shortpath exists error:", err)
		}
	}

	return len(ret) != 0
}

func (dbs *DBService) InsertShortpath(longurl, shortpath string) {
	stmt, err := dbs.DB.Prepare("INSERT INTO url SET longurl=?," +
		"shortpath=?, created_time=?")
	defer stmt.Close()
	if err != nil {
		logq.Fatal("Insert into database error: ", err)
	}

	res, err := stmt.Exec(longurl, shortpath, time.Now())
	if err != nil {
		logq.Fatal(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		logq.Fatal("insert into url error: ", err)
	}
}

func (dbs *DBService) QueryUrlRecord(shortpath string) string {
	stmt, err := dbs.DB.Prepare("SELECT id, longurl, shortpath FROM url WHERE shortpath=?")
	defer stmt.Close()
	if err != nil {
		logq.Fatal("query shortpath record error: ", err)
	}

	row := stmt.QueryRow(shortpath)
	var url Url
	err = row.Scan(&url.Id, &url.Longurl, &url.Shortpath)
	if err != nil {
		logq.Warn("query url records error: ", err)
		return ""
	}

	return url.Longurl
}
