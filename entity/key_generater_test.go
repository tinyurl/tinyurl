package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var indexes = []struct {
	Index    int64
	Expected string
}{
	{0, "A"},
	{1, "B"},
	{25, "Z"},
	{32, "g"},
	{59, "7"},
	{62, "AB"},
	{2016, "gg"},
	{2031, "vg"},
	{131071, "DGi"}, // 2<<16-1
	{131033, "bFi"},
	{77777, "dOU"},
	{12906883, "hpJ2"},
	{12906879, "dpJ2"},
	{12906857, "HpJ2"},
	{12906764, "mnJ2"},
	{12906604, "ClJ2"},
}

var sender = SenderWorker{}

func TestDefaultSenderWorker(t *testing.T) {
	s := DefaultSenderWorker()
	assert.NotEqual(t, s.Index, SenderDefaultIndex)
}

func TestNewSenderWorker(t *testing.T) {
	s := NewSenderWorker(4)
	assert.NotEqual(t, s.Index, 3)
}

func TestSenderNew(t *testing.T) {
	for s := 1; s < 2<<12; s++ {
		ret := sender.New()
		assert.NotNil(t, sender.Index)
		assert.NotNil(t, ret)
	}
}

func TestGetByteByIndex(t *testing.T) {
	for _, I := range indexes {
		ret := GetByteByIndex(I.Index, DefaultChars)
		assert.Equal(t, ret, I.Expected)
	}
}

var basicGenerater = BasicGenerater{}

func TestBasicNew(t *testing.T) {
	u0 := basicGenerater.New()
	assert.Equal(t, len(u0), BasicDefaultLen)
}

func TestBasicNewLen(t *testing.T) {
	u0 := basicGenerater.NewLen(8)
	assert.Equal(t, len(u0), 8)
}

func TestBasicNewLenChars(t *testing.T) {
	u0 := basicGenerater.NewLenChars(4, []byte("123456789"))
	assert.Equal(t, len(u0), 4)
}

func BenchmarkBasicNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		basicGenerater.New()
	}
}

func BenchmarkBasicNewLen(b *testing.B) {
	b.Logf("new uri with 10 lenth using BasicNewLen\n")
	for n := 0; n < b.N; n++ {
		basicGenerater.NewLen(10)
	}

	b.Logf("new uri with 20 lenth using BasicNewLen\n")
	for n := 0; n < b.N; n++ {
		basicGenerater.NewLen(20)
	}
}
