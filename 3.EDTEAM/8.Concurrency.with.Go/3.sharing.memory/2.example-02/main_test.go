package main

import "testing"

func Benchmark_fetch_seq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchSequential(urls)
	}
}

func Benchmark_fetch_conc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchConcurrent(urls)
	}
}
func Benchmark_fetch_scp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchConcurrentSCP(urls)
	}
}

// go test -bench=.
