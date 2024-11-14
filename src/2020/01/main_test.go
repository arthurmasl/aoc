package main

import (
	"testing"

	"aoc/src/internal/utils"
)

var (
	lines = utils.GetLines("./input")
	nums  = utils.ConvertToInts(lines)
)

func Benchmark_solution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1(nums)
	}
}

func Benchmark_solution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2(nums)
	}
}

func Benchmark_solution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution3(nums)
	}
}
