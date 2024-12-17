package main

import (
	"aoc/internal/utils"
	"container/list"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	line := utils.GetLines("src/2024/11/input")[0]
	numbers := utils.ConvertToInts(strings.Split(line, " "))

	// stopProfiling := utils.Profile()
	// defer stopProfiling()

	t1 := time.Now()

	items := list.New()
	for _, n := range numbers {
		items.PushBack(n)
	}

	for iter := range 25 {
		t2 := time.Now()

		for e := items.Front(); e != nil; e = e.Next() {
			n := e.Value.(int)

			if n == 0 {
				e.Value = 1
				continue
			}

			if digits := countDigits(n); digits%2 == 0 {
				l, r := splitNumber(n, digits/2)
				e.Value = l
				items.PushFront(r)
				continue
			}

			e.Value = n * 2024
		}

		fmt.Println(iter, time.Since(t2))
	}

	fmt.Printf("answer: %v time: %v\n", items.Len(), time.Since(t1))
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
