package main

import (
	"flag"
	"fmt"
)

var (
	usi     *UrlServiceImpl
	dbname  string
	user    string
	pass    string
	address string
	dbport  string
	port    string
)

func main() {
	flag.StringVar(&dbname, "dbname", "tinyurl", "database name to connect")
	flag.StringVar(&user, "user", "test", "user of database")
	flag.StringVar(&pass, "pass", "test", "pass of database")
	flag.StringVar(&address, "address", "localhost", "address of database")
	flag.StringVar(&dbport, "dbport", "3306", "port of database")
	flag.StringVar(&port, "port", "8877", "port tinyurl bind on")
	flag.Parse()

	fmt.Println("Start init DB")
	dbs := NewDB(dbname, user, pass, address, dbport)
	usi = NewUrlServiceImpl(dbs)
	tinyUrlAPI(":" + port)
}
