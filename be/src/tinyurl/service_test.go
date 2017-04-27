package main

import(
	"testing"
	"fmt"
)

var u *UrlServiceImpl = NewUrlServiceImpl()

func TestShorten(t *testing.T) {
	fmt.Println("Start test Shorten method...")
	ret1 := u.Shorten(4)
	ret2 := u.Shorten(5)
	ret3 := u.Shorten(6)
	if len(ret1) != 4 || len(ret2) != 5 || len(ret3) != 6 {
		t.Error("Can not get correct length shorten url")
	}
}

func BenchmarkShorten(b *testing.B) {
	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 4")
	for i:=0; i<b.N; i++ {
		u.Shorten(4)
	}
	b.StopTimer()

	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 5")
	for i:=0; i<b.N; i++ {
		u.Shorten(5)
	}
	b.StopTimer()

	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 10")
	for i:=0; i<b.N; i++ {
		u.Shorten(10)
	}
	b.StopTimer()

	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 15")
	for i:=0; i<b.N; i++ {
		u.Shorten(15)
	}
}
