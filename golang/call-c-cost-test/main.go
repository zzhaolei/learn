package main

/*
	int rand() { return 42; }
*/
import "C"

import "testing"

func BenchmarkCgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C.rand()
	}
}

func main() {
	testing.Main(func(string, string) (bool, error) {
		return true, nil
	}, nil, []testing.InternalBenchmark{
		{"BenchmarkCgo", BenchmarkCgo},
	}, nil)
}
