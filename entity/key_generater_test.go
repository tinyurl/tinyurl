package entity

import (
	"math"
	"sync"
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

func TestSyncSenderNew(t *testing.T) {
	s := NewSenderWorker(0)
	var i, n int64
	n = 1000
	wg := sync.WaitGroup{}
	for i = 0; i < n; i++ {
		wg.Add(1)
		go func() {
			s.New()
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, n, s.Index, "index should be %d, but get %d \n", n, s.Index)
}

func TestGetByteByIndex(t *testing.T) {
	for _, I := range indexes {
		ret := GetByteByIndex(I.Index, DefaultChars)
		assert.Equal(t, ret, I.Expected)
	}
}

func TestSenderDefaultLen(t *testing.T) {
	len1 := 2
	len2 := 3
	index1 := int64(math.Pow(float64(62), float64(len1-1)))
	index2 := int64(math.Pow(float64(62), float64(len2-1)))
	sender1 := NewSenderWorker(index1)
	sender2 := NewSenderWorker(index2)

	key1 := sender1.New()
	key2 := sender2.New()

	assert.Equal(t, key1, "AB")
	assert.Equal(t, key2, "AAB")

	// fmt.Printf("key1 %s index1 %d, key2 %s index2 %d\n",
	// 	key1, index1, key2, index2)
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
