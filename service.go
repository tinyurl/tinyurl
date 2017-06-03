package main

import (
	"math/rand"
)

const (
	// ValidChars chars to consists of random string
	ValidChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	// DefaultLen default random string length
	DefaultLen = 4
)

// URLServiceImpl main service of shortpath
type URLServiceImpl struct {
	dbs    *DBService
	length int
}

// NewURLServiceImpl return new UrlServiceImpl instance
func NewURLServiceImpl(dbs *DBService) *URLServiceImpl {
	usi := &URLServiceImpl{dbs: dbs, length: DefaultLen}
	return usi
}

// Shorten generate a short, non-repeat path maps to longurl
func (u *URLServiceImpl) Shorten(longurl string, length int) string {
	if length == 0 {
		u.SetLen(DefaultLen)
	} else {
		u.SetLen(length)
	}

	shortpath := generateRandomStr(u.GetLen())
	for u.Seek(shortpath) {
		shortpath = generateRandomStr(u.GetLen())
	}
	u.Put(longurl, shortpath)

	return shortpath
}

// Put save url and shortpath to db
func (u *URLServiceImpl) Put(longurl, shortpath string) {
	u.dbs.InsertShortpath(longurl, shortpath)
}

// GetShortpath get shortpath and url by shortpath from db
func (u *URLServiceImpl) GetShortpath(shortpath string) string {
	return u.dbs.QueryUrlRecord(shortpath)
}

// SetLen set current random string length
func (u *URLServiceImpl) SetLen(length int) {
	u.length = length
}

// GetLen return current random string length
func (u *URLServiceImpl) GetLen() int {
	return u.length
}

// Seek check if path has existed in db
func (u *URLServiceImpl) Seek(path string) bool {
	return u.dbs.CheckPath(path)
}

// generateRandomStr
func generateRandomStr(length int) string {
	shortURL := make([]byte, length)
	seed := make([]byte, length)
	charLength := len(ValidChars)

	_, err := rand.Read(seed)
	if err != nil {
		logq.Fatal(err)
	}

	for k, v := range seed {
		i := int(v)
		shortURL[k] = ValidChars[i%charLength]
	}

	return string(shortURL)
}
