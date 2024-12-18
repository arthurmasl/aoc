package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "23999685"

	var a, b, c int
	var result string
	a, _ = strconv.Atoi(input)

	for range len(input) + 1 {
		b = ((a & 7) ^ 1)
		c = a >> b
		b ^= 5
		a >>= 3
		b ^= c

		result += strconv.Itoa(b & 7)
	}

	fmt.Println("target 2411751503445530")
	fmt.Println("output", result)
	fmt.Println(a, b, c)

	if result == "2411751503445530" {
		panic("done")
	}
}
