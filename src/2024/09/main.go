package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Cell struct {
	id, size int
}

func main() {
	input, _ := os.ReadFile("src/2024/09/example")

	storage := make([]Cell, 0)
	id := 0

	for chunk := range slices.Chunk(input[:len(input)-1], 2) {
		size, _ := strconv.Atoi(string(chunk[0]))
		space := 0
		if len(chunk) == 2 {
			space, _ = strconv.Atoi(string(chunk[1]))
		}
		storage = append(storage, Cell{id, size})
		storage = append(storage, Cell{-1, space})
		id++
	}

	for _, item := range slices.Backward(storage) {
		if item.id == -1 {
			storage = slices.Delete(storage, len(storage)-1, len(storage))
		} else {
			break
		}
	}

	lastIndex := len(storage) - 1
	visitedId := make(map[Cell]bool)

	for {
		var last Cell

		for i, item := range slices.Backward(storage) {
			if item.id != -1 && !visitedId[item] {
				last = item
				lastIndex = i
				visitedId[last] = true
				break
			}
		}

		if last.id == 1 {
			break
		}

		for firstIndex, first := range storage[:lastIndex] {
			if first.id != -1 {
				continue
			}

			if first.size >= last.size {
				diff := first.size - last.size

				storage[firstIndex].id = last.id
				storage[firstIndex].size = last.size
				storage[lastIndex] = Cell{-1, last.size}
				if diff > 0 {
					storage = slices.Insert(storage, firstIndex+1, Cell{-1, diff})
				}
				break
			}
		}

	}

	sum := 0
	index := 0
	for _, n := range storage {
		for range n.size {
			if n.id != -1 {
				sum += n.id * index
			}
			index++
		}
	}

	fmt.Println(sum)
}

func draw(arr []Cell) {
	for _, n := range arr {
		for range n.size {
			if n.id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(n.id)
			}
		}
	}
	fmt.Println()
}
