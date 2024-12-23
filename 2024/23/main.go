package main

import (
	"fmt"
	"slices"
	"strings"

	"aoc/internal/utils"
)

type Graph map[string][]string

func (graph *Graph) isConnected(u, v string) bool {
	for _, neighbor := range (*graph)[u] {
		if neighbor == v {
			return true
		}
	}

	return false
}

func (graph *Graph) find3Cliques() [][]string {
	cliques := make([][]string, 0)
	vertices := make([]string, 0)

	for vertex := range *graph {
		vertices = append(vertices, vertex)
	}

	for i := 0; i < len(vertices); i++ {
		for j := i + 1; j < len(vertices); j++ {
			for k := j + 1; k < len(vertices); k++ {
				v1, v2, v3 := vertices[i], vertices[j], vertices[k]

				if graph.isConnected(v1, v2) && graph.isConnected(v2, v3) &&
					graph.isConnected(v1, v3) {
					cliques = append(cliques, []string{v1, v2, v3})
				}
			}
		}
	}

	return cliques
}

func (graph *Graph) bronKerbosch(r, p, x []string, cliques *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		*cliques = append(*cliques, append([]string{}, r...))
		return
	}

	for i := 0; i < len(p); i++ {
		v := p[i]
		newR := append(r, v)
		newP := []string{}
		newX := []string{}

		for _, u := range p {
			if graph.isConnected(v, u) {
				newP = append(newP, u)
			}
		}

		for _, u := range x {
			if graph.isConnected(v, u) {
				newX = append(newX, u)
			}
		}

		graph.bronKerbosch(newR, newP, newX, cliques)

		p = append(p[:i], p[i+1:]...)
		x = append(x, v)
		i--
	}
}

func (graph *Graph) findMaxCliques() ([][]string, int, [][]string) {
	cliques := make([][]string, 0)
	vertices := make([]string, 0)

	for vertex := range *graph {
		vertices = append(vertices, vertex)
	}

	graph.bronKerbosch([]string{}, vertices, []string{}, &cliques)

	maxSize := 0
	for _, clique := range cliques {
		if len(clique) > maxSize {
			maxSize = len(clique)
		}
	}

	maxCliques := make([][]string, 0)
	for _, clique := range cliques {
		if len(clique) == maxSize {
			fmt.Println("found")
			maxCliques = append(maxCliques, clique)
		}
	}

	return cliques, maxSize, maxCliques
}

func main() {
	lines := utils.GetLines("input")
	graph := make(Graph)

	for _, line := range lines {
		l, r, _ := strings.Cut(line, "-")

		graph[l] = append(graph[l], r)
		graph[r] = append(graph[r], l)
	}

	cliques := graph.find3Cliques()

	count := 0
	for _, clique := range cliques {
		for _, id := range clique {
			if strings.HasPrefix(id, "t") {
				count++
				break
			}
		}
	}

	fmt.Println(count)

	_, _, maxCliques := graph.findMaxCliques()

	// fmt.Println(allCliques)
	// fmt.Println(maxSize)
	// fmt.Println(maxCliques)
	result := strings.Join(slices.Sorted(slices.Values(maxCliques[0])), ",")
	fmt.Println(result)
}
