package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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

const fileName = "example"

// 503576154 bad
func main() {
	blocks := utils.GetLines(fileName, "\n\n")
	registers := getRegisters(blocks[0])
	opcodes := getProgram(blocks[1])

	pointer := 0
	output := ""
	for pointer != len(opcodes) {
		opcode := opcodes[pointer]

		literalOperand := opcodes[pointer+1]
		comboOperand := getOperandValue(literalOperand, registers)

		switch Opcode(opcode) {
		case adv:
			registers[a] = registers[a] / int(math.Pow(2, float64(comboOperand)))
		case bxl:
			registers[b] ^= literalOperand
		case bst:
			registers[b] = comboOperand % 8
		case jnz:
			if registers[a] != 0 {
				pointer = literalOperand
				continue
			}
		case bxc:
			registers[b] ^= registers[c]
		case out:
			output += strconv.Itoa(comboOperand%8) + ","
		case bdv:
			registers[b] = registers[a] / int(math.Pow(2, float64(comboOperand)))
		case cdv:
			registers[c] = registers[a] / int(math.Pow(2, float64(comboOperand)))
		}

		pointer += 2
	}

	result := strings.ReplaceAll(output, ",", "")
	fmt.Println(result)
	fmt.Println(registers[a], registers[b], registers[c])

	if fileName == "example" {
		utils.Assert(result == "4635635210")
	}
}

func getOperandValue(operand int, registers [3]int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers[a]
	case 5:
		return registers[b]
	case 6:
		return registers[c]
	}

	return -1
}

func getRegisters(block string) [3]int {
	registers := [3]int{}

	for i, l := range strings.Split(block, "\n") {
		_, nStr, _ := strings.Cut(l, ": ")
		n, _ := strconv.Atoi(nStr)
		registers[i] = n
	}
	return registers
}

func getProgram(block string) []int {
	_, numbers, _ := strings.Cut(block, ": ")
	opcodes := utils.ConvertToInts(strings.Split(numbers, ","))

	return opcodes
}
