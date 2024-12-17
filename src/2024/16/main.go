package main

import (
	"container/heap"
	"fmt"
	"slices"

	"aoc/src/internal/utils"
)

type Vector struct {
	x, y float64
}

var (
	directions = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	arrows     = []string{"^", ">", "v", "<"}
)

// 453 to low
func main() {
	lines := utils.GetLines("example")

	startPos := Vector{}
	endPos := Vector{}

	distance := make(map[Vector]int)
	parent := make(map[Vector]Vector)
	path := make([]Vector, 0)
	lowestCost := 0

	paths := make(map[int][][]Vector)

	for y, row := range lines {
		for x, col := range row {
			if col == 'S' {
				startPos = Vector{float64(x), float64(y)}
			}
			if col == 'E' {
				endPos = Vector{float64(x), float64(y)}
			}
		}
	}

	distance[startPos] = 0

	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	heap.Push(&priorityQueue, &Item{value: ItemValue{startPos, arrows[1]}})

	for priorityQueue.Len() > 0 {
		element := heap.Pop(&priorityQueue).(*Item)
		currentPos := element.value.pos
		currentDir := element.value.direction
		currentCost := element.cost

		if currentPos == endPos {
			path = make([]Vector, 0)
			for currentPos != startPos {
				path = append(path, currentPos)
				currentPos = parent[currentPos]
			}
			slices.Reverse(path)
			lowestCost = currentCost
			paths[currentCost] = append(paths[currentCost], path)
			fmt.Println(currentCost)
			continue
		}

		for _, neighbor := range getNeighbors(lines, currentPos) {
			neighborPos := neighbor.pos
			neigborDir := neighbor.direction
			newCost := currentCost + 1
			if neigborDir != currentDir {
				newCost += 1000
			}

			if cost, ok := distance[neighborPos]; !ok || newCost <= cost {
				distance[neighborPos] = newCost
				parent[neighborPos] = currentPos
				heap.Push(
					&priorityQueue,
					&Item{value: ItemValue{neighborPos, neigborDir}, cost: newCost},
				)
			}
		}
	}

	for _, path := range paths[lowestCost] {
		// lines := utils.GetLines("example")
		// fill path
		for _, pos := range path {
			for y, row := range lines {
				newRow := []byte(row)
				for x := range row {
					if int(pos.x) == x && int(pos.y) == y {
						newRow[x] = 'O'
					}
				}
				lines[y] = string(newRow)
			}
		}
	}

	// draw grid
	for _, row := range lines {
		fmt.Println(row)
	}
	fmt.Println()

	total := 1
	for _, row := range lines {
		for _, col := range row {
			if col == 'O' {
				total++
			}
		}
	}

	fmt.Println(len(paths[lowestCost]))
	fmt.Println(lowestCost, total)

	fmt.Println(priorityQueue.Len())
}

type Neighbor struct {
	pos       Vector
	direction string
}

func getNeighbors(lines []string, pos Vector) []Neighbor {
	neighbors := make([]Neighbor, 0)

	for i, dir := range directions {
		neighborPos := Vector{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(lines, int(neighborPos.x), int(neighborPos.y))

		if ok && neighborId != '#' {
			neighbors = append(neighbors, Neighbor{neighborPos, arrows[i]})
		}
	}

	return neighbors
}

type ItemValue struct {
	pos       Vector
	direction string
}

type Item struct {
	value ItemValue
	cost  int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value ItemValue, priority int) {
	item.value = value
	item.cost = priority
	heap.Fix(pq, item.index)
}
