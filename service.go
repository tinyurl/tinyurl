package main

import(
	"math/rand"
)

const(
	VALIDECHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	DEFAULTLEN = 4
)

type UrlServiceImpl struct {
	dbs 	*DBService
	length int
}

func NewUrlServiceImpl(dbs *DBService) *UrlServiceImpl {
	usi := &UrlServiceImpl{dbs: dbs, length: DEFAULTLEN,}
	return usi
}

// Shorten generate a short, non-repeat path maps to longurl
func (u *UrlServiceImpl) Shorten(longurl string, length int) string {
	if length == 0{
		u.SetLen(DEFAULTLEN)
	}else{
		u.SetLen(length)
	}

	shortpath := generateRandomStr(u.GetLen())
	for u.Seek(shortpath) {
		shortpath = generateRandomStr(u.GetLen())
	}
	u.Put(longurl, shortpath)

	return shortpath
}

func (u *UrlServiceImpl) Put(longurl, shortpath string) {
	u.dbs.InsertShortpath(longurl, shortpath)
}

func (u *UrlServiceImpl) GetShortpath(shortpath string) string {
	return u.dbs.QueryUrlRecord(shortpath)
}

func (u *UrlServiceImpl) SetLen(length int) {
	u.length = length
}

func (u *UrlServiceImpl) GetLen() int {
	return u.length
}

func generateRandomStr(length int) string {
	var shortUrl []byte 
	for i:=0; i<length; i++ {
		randnum := rand.Intn(len(VALIDECHARS))
		shortUrl = append(shortUrl, VALIDECHARS[randnum])
	}

	return string(shortUrl)
}

// seek check if path has existed in db
func (u *UrlServiceImpl) Seek(path string) bool {
	return u.dbs.CheckPath(path)	
}
