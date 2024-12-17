package main

import (
	"os"
	"testing"
)

var (
	line, _ = os.ReadFile("./input.txt")
	str     = line[:len(line)-1]
)

func Benchmark_solution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1(str)
	}
}

func Benchmark_solution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2(str)
	}
}

func Benchmark_solution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution3(str)
	}
}

func Benchmark_solution4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution4(str)
	}
}

func Benchmark_solution5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution5(str)
	}
}
