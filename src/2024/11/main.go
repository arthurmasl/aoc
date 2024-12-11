package main

import (
	"aoc/src/internal/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	line := utils.GetLines("src/2024/11/input")[0]
	numbers := utils.ConvertToInts(strings.Split(line, " "))

	t1 := time.Now()
	memo := make(map[int]int)

	for _, n := range numbers {
		memo[n] += 1
	}

	for range 75 {
		newNumbers := make(map[int]int)

		for val, count := range memo {
			if val == 0 {
				newNumbers[1] += count
				continue
			}

			if countDigits(val)%2 == 0 {
				l, r := splitNumber(val)
				newNumbers[l] += count
				newNumbers[r] += count
				continue
			}

			newNumbers[val*2024] += count
		}

		memo = newNumbers
	}

	total := 0
	for _, v := range memo {
		total += v
	}

	fmt.Println(time.Since(t1))
	fmt.Println(total)
}

func countDigits(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}

func splitNumber(n int) (int, int) {
	str := strconv.Itoa(n)
	half := len(str) / 2

	l, r := str[:half], str[half:]
	left, _ := strconv.Atoi(l)
	right, _ := strconv.Atoi(r)

	return left, right
}
