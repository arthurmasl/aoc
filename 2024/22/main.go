package main

import (
	"fmt"
	"slices"

	"aoc/internal/utils"
)

func main() {
	numbers := utils.ConvertToInts(utils.GetLines("example"))

	targetSeq := []int{-2, 1, -1, 3}
	seq := make([]int, 4)

	sum := 0

	for _, n := range numbers {
		secret := n
		price := secret % 10

		for range 2000 {
			newSecret := generate(secret, 1)
			newPrice := newSecret % 10
			diff := newPrice - price

			secret = newSecret
			price = newPrice
			seq = seq[1:]
			seq = append(seq, diff)

			if slices.Equal(seq, targetSeq) {
				sum += price
				break
			}
		}
	}

	fmt.Println(sum)
}

func generate(s, n int) int {
	for range n {
		s = prune(mix(s, s*64))
		s = prune(mix(s, s/32))
		s = prune(mix(s, s*2048))
	}

	return s
}

func mix(s, n int) int {
	return n ^ s
}

func prune(s int) int {
	return s % 16777216
}
