package entity

import (
	"crypto/rand"
	"log"
)

var (
	// DefaultChars chars to consists of random string
	DefaultChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

const (
	// BasicDefaultLen default random string length
	BasicDefaultLen    = 8
	SenderDefaultStart = 0

	KeyAlgoRandom = "random"
	KeyAlgoSender = "sender"
)

// KeyGenerater interface for extension
type KeyGenerater interface {
	New() string
	NewLen(int) string
	NewLenChars(int, []byte) string
}

// NewKeyGenerater return KeyGenerater by algoritm
func NewKeyGenerater(algo string) KeyGenerater {
	switch algo {
	case KeyAlgoRandom:
		return BasicGenerater{}
	case KeyAlgoSender:
		return DefaultSenderGenerater()
	default:
		return BasicGenerater{}
	}
}

type SenderGenerater struct {
	Start int64
}

// DefaultSenderGenerater return default SenderGenerater
func DefaultSenderGenerater() *SenderGenerater {
	return NewSenderGenerater(SenderDefaultStart)
}

// NewSenderGenerater
func NewSenderGenerater(start int64) *SenderGenerater {
	sender := &SenderGenerater{
		Start: start,
	}

	return sender
}

// New new uri with default length and []byte
func (sender *SenderGenerater) New() string {
	// m.RLock()
	// defer m.RUnlock()
	key := GetByteByIndex(sender.Start, DefaultChars)
	sender.Start++
	return key
}

// New new key with default length and []byte
func (sender *SenderGenerater) NewLen(length int) string {
	// m.RLock()
	// defer m.RUnlock()
	key := GetByteByIndex(sender.Start, DefaultChars)
	sender.Start++
	return key
}

// New new key with default length and []byte
func (sender *SenderGenerater) NewLenDefaultChars(length int, DefaultChars []byte) string {
	// m.RLock()
	// defer m.RUnlock()
	key := GetByteByIndex(sender.Start, DefaultChars)
	sender.Start++
	return key
}

func (sender *SenderGenerater) NewLenChars(length int, chars []byte) string {
	key := GetByteByIndex(sender.Start, DefaultChars)
	sender.Start++
	return key
}

func GetByteByIndex(index int64, DefaultChars []byte) string {
	baseLen := int64(len(DefaultChars))
	var container []byte
	for ; index != 0; index = index / baseLen {
		m := index % baseLen
		container = append(container, DefaultChars[m])
	}

	return string(container)
}

type BasicGenerater struct {
	KeyLen int
}

// New new BasicGenerater with default length and []byte
func (basic BasicGenerater) New() string {
	return basic.NewLenChars(BasicDefaultLen, DefaultChars)
}

// NewLen new BasicGenerater with given length and default []byte
func (basic BasicGenerater) NewLen(length int) string {
	return basic.NewLenChars(length, DefaultChars)
}

// NewLenChars new BasicGenerater with given lenth and []byte
func (basic BasicGenerater) NewLenChars(length int, chars []byte) string {
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
