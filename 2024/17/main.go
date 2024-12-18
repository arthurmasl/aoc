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
	output := toInt(program(input))

	fmt.Println("input ", input)
	fmt.Println("output", output)
	fmt.Println("target", target)
	fmt.Println()

	utils.Assert(toInt(program(exampleInput)) == exampleOutput)
	// utils.Assert(toInt(program(reverseProgram(target))) == target)
}

func program(a int) []int {
	var result []int

	for a != 0 {
		b := ((a & 7) ^ 1)
		c := a >> b
		b ^= 5
		a >>= 3
		b ^= c

		output := b & 7
		result = append(result, output)
	}

	return result
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

func toInt(ints []int) int {
	str := ""
	for _, n := range ints {
		str += strconv.Itoa(n)
	}
	number, _ := strconv.Atoi(str)
	return number
}
