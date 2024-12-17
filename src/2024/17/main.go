package main

import (
	"fmt"
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

func main() {
	blocks := utils.GetLines(fileName, "\n\n")
	registers := getRegisters(blocks[0])
	opcodes := getProgram(blocks[1])

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
			registers[b] ^= literalOperand
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
			output += strconv.Itoa(comboValue&7) + ","
		case bdv:
			registers[b] = registers[a] >> comboValue
		case cdv:
			registers[c] = registers[a] >> comboValue
		}

		pointer += 2
	}

	result := output[:len(output)-1]
	fmt.Println(result)
	// fmt.Println(registers[a], registers[b], registers[c])

	if fileName == "example" {
		utils.Assert(result == "4,6,3,5,6,3,5,2,1,0")
	}
}

func getComboValue(operand int, registers [3]int) int {
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
	_, numsStr, _ := strings.Cut(block, ": ")
	return utils.ConvertToInts(strings.Split(numsStr, ","))
}
