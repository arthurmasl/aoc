package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc/src/internal/utils"
)

type Register int

const (
	a Register = iota
	b
	c
)

type Opcode int

const (
	adv Opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

const fileName = "input"

func main() {
	fromG := int64(139100100000181)
	// to := 290000000000000
	// r := to - from

	ranges := [12][2]int64{
		{1, 12574991666651},
		{12574991666652, 25149983333302},
		{25149983333303, 37724974999953},
		{37724974999954, 50299966666604},
		{50299966666605, 62874958333255},
		{62874958333256, 75449949999906},
		{75449949999907, 88024941666557},
		{88024941666558, 100599933333208},
		{100599933333209, 113174924999859},
		{113174924999860, 125749916666510},
		{125749916666511, 138324908333161},
		{138324908333162, 150899899999819},
	}

	for i, chunk := range ranges {
		from := chunk[0]
		to := chunk[1]
		r := to - from

		go func(threadId int) {
			fmt.Println("start", threadId)

			blocks := utils.GetLines(fileName, "\n\n")
			registers := getRegisters(blocks[0])
			opcodes := getProgram(blocks[1])

			for i := range r {
				registers[a] = from + fromG + i

				pointer := 0
				output := ""

				for pointer != len(opcodes) {
					opcode := opcodes[pointer]
					literalOperand := int(opcodes[pointer+1])
					comboValue := getComboValue(literalOperand, registers)

					switch Opcode(opcode) {
					case adv:
						registers[a] >>= comboValue
					case bxl:
						registers[b] ^= int64(literalOperand)
					case bst:
						registers[b] = comboValue & 7
					case jnz:
						if registers[a] != 0 {
							pointer = int(literalOperand)
							continue
						}
					case bxc:
						registers[b] ^= registers[c]
					case out:
						output += strconv.Itoa(int(comboValue&7)) + ","
					case bdv:
						registers[b] = registers[a] >> comboValue
					case cdv:
						registers[c] = registers[a] >> comboValue
					}

					pointer += 2
				}

				// from 35000000000000
				// to 290000000000000

				result := output[:len(output)-1]
				// fmt.Println(len(result)/2 + 1)
				// fmt.Println(result)
				// fmt.Println(registers[a], registers[b], registers[c])

				if i%10000000 == 0 {
					fmt.Printf("thread %v, result: %v (%v)\n", threadId, result, len(result)/2+1)
				}

				if result == "2,4,1,1,7,5,1,5,0,3,4,4,5,5,3,0" {
					fmt.Println("===found", from+i)
					break
				}
			}
		}(i)
	}

	time.Sleep(time.Hour * 2)
}

func getComboValue(operand int, registers [3]int64) int64 {
	switch operand {
	case 0, 1, 2, 3:
		return int64(operand)
	case 4:
		return registers[a]
	case 5:
		return registers[b]
	case 6:
		return registers[c]
	}

	return -1
}

func getRegisters(block string) [3]int64 {
	registers := [3]int64{}

	for i, l := range strings.Split(block, "\n") {
		_, nStr, _ := strings.Cut(l, ": ")
		n, _ := strconv.Atoi(nStr)
		registers[i] = int64(n)
	}
	return registers
}

func getProgram(block string) []int {
	_, numsStr, _ := strings.Cut(block, ": ")
	return utils.ConvertToInts(strings.Split(numsStr, ","))
}
