package main

import "fmt"

func main() {
	a := 23999685
	for range 16 {
		b := ((a & 7) ^ 1) ^ 5
		c := a >> 4
		a >>= 3
		b ^= c

		fmt.Print(b&7, ", ")
	}

	fmt.Println()
}
