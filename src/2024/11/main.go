package main

import (
	"aoc/src/internal/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	line := utils.GetLines("src/2024/11/input")[0]
	numbers := utils.ConvertToInts(strings.Split(line, " "))
	// workers := runtime.GOMAXPROCS(runtime.NumCPU())

	stopProfiling := utils.Profile()
	defer stopProfiling()

	fmt.Println(numbers)
	t1 := time.Now()

	// 25 - 3ms
	for iteration := range 25 {
		t2 := time.Now()

		for i, n := range numbers {
			if n == 0 {
				numbers[i] = 1
				continue
			}

			if digits := countDigits(n); digits%2 == 0 {
				l, r := splitNumber(n, digits/2)
				numbers[i] = l
				numbers = append(numbers, r)
				continue
			}

			numbers[i] = n * 2024
		}

		fmt.Println(iteration, time.Since(t2))
	}

	fmt.Printf("answer: %v time: %v\n", len(numbers), time.Since(t1))
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func splitNumber(number, position int) (int, int) {
	divisor := int(math.Pow10(position))

	left := number / divisor
	right := number % divisor

	return left, right
}
