package entity

import (
	"crypto/rand"
	"log"
	"sync"
)

var (
	// DefaultChars chars to consists of random string
	DefaultChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

const (
	// BasicDefaultLen default random string length
	BasicDefaultLen    = 8
	SenderDefaultIndex = 0

	KeyAlgoRandom = "random"
	KeyAlgoSender = "sender"
)

// KeyGenerater interface for extension
type KeyGenerater interface {
	New() string
	NewLen(int) string
	NewLenChars(int, []byte) string
	SetIndex(index int64)
	GetIndex() int64
}

// NewKeyGenerater return KeyGenerater by algoritm
func NewKeyGenerater(algo string) KeyGenerater {
	switch algo {
	case KeyAlgoRandom:
		return BasicGenerater{}
	case KeyAlgoSender:
		return DefaultSenderWorker()
	default:
		return BasicGenerater{}
	}
}

// SenderManager for concurrent process
// type SenderManager struct {
// 	NumWorker int
// 	Step      int
// 	Config    map[int]*SenderWorker
// }

type SenderWorker struct {
	// TODOs: adjust for multi db
	ID    int64      `gorm:"primary_key"`
	Index int64      `gorm:"type:int"`
	m     sync.Mutex `gorm:"-"` // ignore
}

// DefaultSenderWorker return default SenderWorker
func DefaultSenderWorker() *SenderWorker {
	return NewSenderWorker(SenderDefaultIndex)
}

// NewSenderWorker
func NewSenderWorker(index int64) *SenderWorker {
	sender := &SenderWorker{
		Index: index,
		m:     sync.Mutex{},
	}

	return sender
}

func (sender *SenderWorker) GetIndex() int64 {
	return sender.Index
}

// SetIndex
func (sender *SenderWorker) SetIndex(index int64) {
	sender.Index = index
}

// New new uri with default length and []byte
func (sender *SenderWorker) New() string {
	sender.m.Lock()
	defer sender.m.Unlock()
	key := GetByteByIndex(sender.Index, DefaultChars)
	sender.Index++
	return key
}

// New new key with default length and []byte
func (sender *SenderWorker) NewLen(length int) string {
	// this function just for interface
	return ""
}

func (sender *SenderWorker) NewLenChars(length int, chars []byte) string {
	// this function just for interface
	return ""
}

func GetByteByIndex(index int64, DefaultChars []byte) string {
	baseLen := int64(len(DefaultChars))
	if index == 0 {
		return string(DefaultChars[0])
	}
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

func (basic BasicGenerater) GetIndex() int64 {
	return 0
}

func (basic BasicGenerater) SetIndex(index int64) {

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
