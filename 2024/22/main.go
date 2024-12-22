package main

import (
	"fmt"

	"aoc/internal/utils"
)

type stat struct {
	diff  int
	price int
}

type seq [4]int

func main() {
	numbers := utils.ConvertToInts(utils.GetLines("input"))
	total := make(map[seq]int)

	for _, n := range numbers {
		memo := make(map[seq]bool)
		stats := make([]stat, 0, 2000)

		secret := n
		price := secret % 10

		for range 2000 {
			secret = generate(secret)
			newPrice := secret % 10
			diff := newPrice - price
			price = newPrice

			stats = append(stats, stat{diff, price})
		}

		for statChunk := range utils.Window(stats, 4) {
			newSeq := seq{}
			for i, newStat := range statChunk {
				newSeq[i] = newStat.diff
			}

			price := statChunk[len(statChunk)-1].price

			if !memo[newSeq] {
				memo[newSeq] = true
				total[newSeq] += price
			}

		}
	}

	largest := 0
	for _, v := range total {
		if v > largest {
			largest = v
		}
	}

	fmt.Println(largest)
}

func generate(s int) int {
	newSecret := prune(mix(s, s*64))
	newSecret = prune(mix(newSecret, newSecret/32))
	newSecret = prune(mix(newSecret, newSecret*2048))

	return newSecret
}

func mix(s, n int) int {
	return n ^ s
}

func prune(s int) int {
	return s % 16777216
}
