package main

import "testing"

func BenchmarkPartA(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		partA()
	}
}

func BenchmarkPartB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		partB()
	}
}
