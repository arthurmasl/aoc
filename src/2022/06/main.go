package main

import (
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	line, _ := os.ReadFile("src/2022/06/input.txt")
	str := line[:len(line)-1]

	t1 := time.Now()
	fmt.Println(solution5(str))
	fmt.Println(time.Since(t1))
}

// 1.5-3ms
func solution1(str []byte) int {
	index := 0
	chunk := make(map[string]bool)

	for {
		for _, c := range str[index : index+14] {
			chunk[string(c)] = true
		}

		if len(chunk) == 14 {
			return index + 14
		}

		clear(chunk)
		index++
	}
}

// 400us
func solution2(str []byte) int {
	index := 0
	chunk := make(map[byte]bool, 14)

LL:
	for {
		for _, c := range str[index : index+14] {
			if ok := chunk[c]; ok {
				clear(chunk)
				index++

				continue LL
			}

			chunk[c] = true
		}

		return index + 14
	}
}

// 100us
func solution3(str []byte) int {
	index := 0
	chunk := make([]byte, 14)

LL:
	for {
		for i, c := range str[index : index+14] {
			for _, b := range chunk {
				if b == c {
					clear(chunk)
					index++
					continue LL
				}
			}

			chunk[i] = c
		}

		return index + 14
	}
}

// 30us
func solution4(str []byte) int {
	index := 0
	gi := 0
	chunk := make([]byte, 14)

LL:
	for {
		for i, c := range str[index : index+14] {
			if slices.Contains(chunk, c) {
				clear(chunk)
				gi++
				index = gi
				continue LL
			}

			chunk[i] = c
			gi++
		}

		return index + 14
	}
}

// 15us
func solution5(str []byte) int {
	index := 0
	gi := 0
	state := uint32(0)

LL:
	for {
		for _, c := range str[index : index+14] {
			bitIdx := (c % 32)

			if state&(1<<bitIdx) != 0 {
				state = 0
				index = gi + 1
				gi++
				continue LL
			}

			state |= 1 << (bitIdx)
			gi++
		}

		return index + 14
	}
}
