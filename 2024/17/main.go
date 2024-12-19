package main

import (
	"fmt"
	"strconv"

	"aoc/internal/utils"
)

const (
	exampleInput  = 23999685
	exampleOutput = "503576154"
	target        = "2411751503445530"
)

func main() {
	a := reverse(target, 0)
	res := program(a)

	fmt.Println(a)
	fmt.Println(res)

	utils.Assert(res == target)
	// utils.Assert(program(exampleInput) == exampleOutput)
}

func program(a int) string {
	var result string

	// fmt.Println("=============================")
	// fmt.Printf("A   == %v\n\n", a)

	for a > 0 {
		b := a%8 ^ 1
		// fmt.Printf("B   -> A mod 8 ^ 1 (%v)\n", b)

		c := a >> b
		// fmt.Printf("C   -> A >> B      (%v)\n", c)

		b ^= 5 ^ c
		// fmt.Printf("B   -> B ^ 5 ^ C   (%v)\n", b)

		a >>= 3
		// fmt.Printf("A   -> A >> 3      (%v)\n", a)

		output := b % 8
		// fmt.Printf("OUT -> B mod 8     (%v)\n\n", output)

		result += strconv.Itoa(output)
	}

	// fmt.Printf("OUT     == %v\n", result[:len(result)-2])
	return result
}

func reverse(input string, a int) int {
	if len(input) == 0 {
		return a
	}

	for i := range 8 {
		a := (a << 3) + i
		b := a%8 ^ 1
		c := a >> b
		b ^= 5 ^ c

		output := b % 8

		s := len(input) - 1
		outputStr := strconv.Itoa(output)
		lastInputStr := string(input[s])

		if outputStr == lastInputStr {
			sub := reverse(input[:s], a)
			if sub == -1 {
				continue
			}
			return sub
		}
	}

	return -1
}
