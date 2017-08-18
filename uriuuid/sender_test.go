package uriuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var indexes = []struct {
	Index    int64
	Expected string
}{
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

var senderURIUUID = SenderURIUUID{}

func TestSenderNew(t *testing.T) {
	for s := 1; s < 2<<12; s++ {
		ret := senderURIUUID.New()
		assert.NotNil(t, senderURIUUID.start)
		assert.NotNil(t, ret)
	}
}

func TestGetByteByIndex(t *testing.T) {
	for _, I := range indexes {
		ret := GetByteByIndex(I.Index, DefaultChars)
		assert.Equal(t, ret, I.Expected)
	}
}
