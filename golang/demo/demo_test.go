package main

import (
	"math/rand"
	"testing"
)

func bit() {
	var s uint64 = uint64(rand.Int63n(1000000000000000000))
	v := int32(s >> 32)
	_ = uint32(v)
}

func notBit() {
	var s uint64 = uint64(rand.Int63n(1000000000000000000))
	_ = uint32(s)
}

func BenchmarkBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bit()
	}
}

func BenchmarkNoBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		notBit()
	}
}
