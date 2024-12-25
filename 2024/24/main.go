package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	blocks := utils.GetLines("input", "\n\n")
	variablesStr := strings.Split(blocks[0], "\n")
	operations := strings.Split(blocks[1], "\n")

	variables := make(map[string]string)

	for _, v := range variablesStr {
		key, valueStr, _ := strings.Cut(v, ": ")
		variables[key] = valueStr
	}

	replaceVariables := func() {
		for i := range operations {
			for k, v := range variables {
				operations[i] = strings.ReplaceAll(operations[i], k, v)
			}
		}
	}

	processOperations := func() {
		for _, line := range operations {
			chunks := strings.Split(line, " ")
			v1 := chunks[0]
			op := chunks[1]
			v2 := chunks[2]
			key := chunks[4]

			if (v1 == "1" || v1 == "0") && (v2 == "1" || v2 == "0") && (key != "1" || key != "0") {
				vv1, _ := strconv.Atoi(v1)
				vv2, _ := strconv.Atoi(v2)

				res := 0
				switch op {
				case "AND":
					res = vv1 & vv2
				case "OR":
					res = vv1 | vv2
				case "XOR":
					res = vv1 ^ vv2

				}
				variables[key] = strconv.Itoa(res)
			}
		}
	}

	for range 100 {
		replaceVariables()
		processOperations()
	}

	fmt.Println(variables)
	fmt.Println()
	for _, line := range operations {
		fmt.Println(line)
	}

	bits := make([]byte, 100)
	for k, v := range variables {
		if strings.HasPrefix(k, "z") {
			indexStr := strings.ReplaceAll(k, "z", "")
			index, _ := strconv.Atoi(indexStr)
			bits[index] = v[0]
		}
	}

	slices.Reverse(bits)
	binaryStr := strings.ReplaceAll(string(bits), "\x00", "")
	dec, _ := strconv.ParseInt(binaryStr, 2, 64)
	fmt.Println(binaryStr)
	fmt.Println(dec)
}
