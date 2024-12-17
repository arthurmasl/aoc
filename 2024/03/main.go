package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("src/2024/03/input")
	match, _ := regexp.Compile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)

	matches := match.FindAllString(string(input), 9999)
	fmt.Println(matches)

	do := true
	res := 0
	for _, match := range matches {
		if match == "do()" {
			do = true
			continue
		}

		if match == "don't()" {
			do = false
			continue
		}

		if do {
			aStr, bStr, _ := strings.Cut(match[4:len(match)-1], ",")
			a, _ := strconv.Atoi(aStr)
			b, _ := strconv.Atoi(bStr)
			res += a * b
		}
	}

	fmt.Println(res)
}
