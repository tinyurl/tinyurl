package uriuuid

import (
	"crypto/rand"
	"log"
)

// BasicURIUUID basic implement of UriUUID interface
type BasicURIUUID struct{}

// New new uri with default length and []byte
func (basic BasicURIUUID) New() string {
	return basic.NewLenChars(DefaultLen, DefaultChars)
}

// NewLen new uri with given length and default []byte
func (basic BasicURIUUID) NewLen(length int) string {
	return basic.NewLenChars(length, DefaultChars)
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
