package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomStr(t *testing.T) {
	fmt.Println("Start test Shorten method...")
	ret1 := generateRandomStr(4)
	ret2 := generateRandomStr(5)
	ret3 := generateRandomStr(6)
	if len(ret1) != 4 || len(ret2) != 5 || len(ret3) != 6 {
		t.Error("Can not get correct length shorten url")
	}

	fmt.Println(ret1, ret2, ret3)
}

func BenchmarkGenerateRandomStr(b *testing.B) {
	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 4")
	for i := 0; i < b.N; i++ {
		generateRandomStr(4)
	}
	b.StopTimer()

	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 5")
	for i := 0; i < b.N; i++ {
		generateRandomStr(5)
	}
	b.StopTimer()

	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 10")
	for i := 0; i < b.N; i++ {
		generateRandomStr(10)
	}
	b.StopTimer()

	b.StartTimer()
	fmt.Println("Start benchmark Shorten with length 15")
	for i := 0; i < b.N; i++ {
		generateRandomStr(15)
	}
}
