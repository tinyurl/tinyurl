package uriuuid

import (
	"crypto/rand"
	"log"
)

var (
	// BasicChars chars to consists of random string
	BasicChars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	// DefaultLen default random string length
	BasicLen = 6
)

// BasicURIUUID basic implement of UriUUID interface
type BasicURIUUID struct{}

// New new uri with default length and []byte
func (basic BasicURIUUID) New() string {
	return basic.NewLenChars(BasicLen, BasicChars)
}

// NewLen new uri with given length and default []byte
func (basic BasicURIUUID) NewLen(length int) string {
	return basic.NewLenChars(length, BasicChars)
}

// NewLenChars new uri with given lenth and []byte
func (basic BasicURIUUID) NewLenChars(length int, chars []byte) string {
	charsLen := len(chars)
	uri := make([]byte, length)
	uriIndex := make([]byte, length)

	_, err := rand.Read(uriIndex)
	if err != nil {
		log.Fatalln("generate random index for chars error: ", err)
	}

	for k, v := range uriIndex {
		index := int(v)
		uri[k] = chars[index%charsLen]
	}

	return string(uri)
}
