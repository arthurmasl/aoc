package main

import (
	"fmt"
	"time"

	"aoc/internal/utils"
)

// 1498 to high
func main() {
	numbers := utils.ConvertToInts(utils.GetLines("example"))

	t1 := time.Now()
	seqs := make(map[[4]int]bool)

	for _, n := range numbers {
		diffs := make([]int, 2000)

		secret := n
		price := secret % 10

		for i := range 2000 {
			newSecret := generate(secret)
			newPrice := newSecret % 10
			diff := newPrice - price

			secret = newSecret
			price = newPrice

			diffs[i] = diff
		}

		for seq := range utils.Window(diffs, 4) {
			arr := [4]int{seq[0], seq[1], seq[2], seq[3]}
			seqs[arr] = true
		}
	}

	largest := 0
	i := 0
	for targetSeq := range seqs {
		i++
		seq := [4]int{}
		sum := 0

		for _, n := range numbers {
			secret := n
			price := secret % 10

			for range 2000 {
				newSecret := generate(secret)
				newPrice := newSecret % 10
				diff := newPrice - price

				secret = newSecret
				price = newPrice

				seq[0] = seq[1]
				seq[1] = seq[2]
				seq[2] = seq[3]
				seq[3] = diff

				if targetSeq == seq {
					sum += price
					break
				}
			}
		}

		if sum > largest {
			fmt.Println(targetSeq, sum)
			largest = sum
		}
	}

	fmt.Println(largest, time.Since(t1))
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
