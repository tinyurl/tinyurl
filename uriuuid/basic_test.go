package uriuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var basicURIUUID = BasicURIUUID{}

func TestBasicNew(t *testing.T) {
	u0 := basicURIUUID.New()
	assert.Equal(t, len(u0), BasicLen)
}

func TestBasicNewLen(t *testing.T) {
	u0 := basicURIUUID.NewLen(8)
	assert.Equal(t, len(u0), 8)
}

func TestBasicNewLenChars(t *testing.T) {
	u0 := basicURIUUID.NewLenChars(4, []byte("123456789"))
	assert.Equal(t, len(u0), 4)
}

func BenchmarkBasicNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		basicURIUUID.New()
	}
}

func BenchmarkBasicNewLen(b *testing.B) {
	b.Logf("new uri with 10 lenth using BasicNewLen\n")
	for n := 0; n < b.N; n++ {
		basicURIUUID.NewLen(10)
	}

	b.Logf("new uri with 20 lenth using BasicNewLen\n")
	for n := 0; n < b.N; n++ {
		basicURIUUID.NewLen(20)
	}
}
