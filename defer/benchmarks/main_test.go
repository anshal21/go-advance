package main

import "testing"

func Benchmark_FuncWithoutDefer(b *testing.B) {
	
	for i := 0; i < b.N; i++ {
		FuncWithoutDefer()
	}
}

func Benchmark_FuncWithDefer(b *testing.B) {

	for i := 0; i < b.N; i++ {
		FuncWithDefer()
	}

}

