package main

import (
	"fmt"
	"strconv"

	"aoc/internal/utils"
)

const (
	exampleInput  = 23999685
	exampleOutput = 503576154
	target        = 2411751503445530
)

func main() {
	input := exampleInput
	output := program(input)

	fmt.Println("input ", input)
	fmt.Println("output", output)
	fmt.Println("target", target)
	fmt.Println()

	utils.Assert(program(exampleInput) == exampleOutput)
	utils.Assert(program(reverseProgram(target)) == target)
}

func program(input int) int {
	var a, b, c int
	var result string

	a = input
	digits := len(strconv.Itoa(input)) + 1

	for range digits {
		b = ((a & 7) ^ 1)
		c = a >> b
		b ^= 5
		a >>= 3
		b ^= c

		result += strconv.Itoa(b & 7)
	}

	res, _ := strconv.Atoi(result)
	return res
}

func reverseProgram(target int) int {
	targetStr := strconv.Itoa(target)
	digits := len(targetStr)

	var a, b, c int

	for i := digits - 1; i >= 0; i-- {
		b, _ = strconv.Atoi(string(targetStr[i]))

		c = (a >> (b ^ 1)) & 7
		b ^= c
		b ^= 5

		a <<= 3
		a |= (b & 7)
	}

	return a
}
