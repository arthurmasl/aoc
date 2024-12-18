package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/internal/utils"
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
	// blocks := utils.GetLines(fileName, "\n\n")
	// registers := getRegisters(blocks[0])
	// opcodes := getProgram(blocks[1])
	//
	// fmt.Println(blocks[1])
	//
	// pointer := 0
	// output := ""
	//
	// operations := []string{}
	//
	// for pointer != len(opcodes) {
	// 	opcode := opcodes[pointer]
	// 	literalOperand := int(opcodes[pointer+1])
	// 	comboValue := getComboValue(literalOperand, registers)
	//
	// 	switch Opcode(opcode) {
	// 	case adv:
	// 		registers[a] >>= comboValue
	// 		operations = append(operations, fmt.Sprintf("a>>=%v=%v ", comboValue, registers[a]))
	// 	case bxl:
	// 		registers[b] ^= int64(literalOperand)
	// 		operations = append(
	// 			operations,
	// 			fmt.Sprintf("b^=%v=%v ", literalOperand, registers[b]))
	// 	case bst:
	// 		registers[b] = comboValue & 7
	// 		operations = append(operations, fmt.Sprintf("b=%-10v&7=%v ", comboValue, comboValue&7))
	// 	case jnz:
	// 		if registers[a] != 0 {
	// 			pointer = int(literalOperand)
	// 			// operations = append(operations, fmt.Sprint("jnz ", literalOperand))
	// 			operations = append(operations, fmt.Sprintf("jmp %v\n", literalOperand))
	// 			continue
	// 		}
	// 	case bxc:
	// 		registers[b] ^= registers[c]
	// 		operations = append(
	// 			operations,
	// 			fmt.Sprintf("b^=c(%v)=%v ", registers[b], registers[b]),
	// 		)
	// 	case out:
	// 		output += strconv.Itoa(int(comboValue&7)) + ","
	// 		operations = append(operations, fmt.Sprintf("out (%v&7)=%v ", comboValue, comboValue&7))
	// 	case bdv:
	// 		registers[b] = registers[a] >> comboValue
	// 		operations = append(operations, "bdv")
	// 	case cdv:
	// 		registers[c] = registers[a] >> comboValue
	// 		operations = append(
	// 			operations,
	// 			fmt.Sprintf("c=a>>%v=%v ", comboValue, registers[c]),
	// 		)
	// 	}
	//
	// 	pointer += 2
	// }

	// result := output[:len(output)-1]
	// fmt.Println("Output: ", result)
	// fmt.Println()
	// fmt.Println(strings.Join(operations, ""))

	// fmt.Println()
	smol()
}

func smol() {
	a := 23999685
	result := ""

	for range 9 {
		b := ((a & 7) ^ 1)
		c := a >> b
		b ^= 5
		a >>= 3
		b ^= c

		result += strconv.Itoa(b & 7)
	}

	fmt.Println("target 2411751503445530")
	fmt.Println("output", result)

	if result == "2411751503445530" {
		panic("done")
	}
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

// func bruteforce() {
// 	fromG := 139100100000181
// 	toG := 290000000000000
//
// 	delta := (toG - fromG) / 11
//
// 	for threadId := range 12 {
// 		from := fromG + delta*threadId
// 		to := from + delta
//
// 		go func(threadId int) {
// 			fmt.Println("Start thread", threadId)
//
// 			for i := from; i < to; i++ {
// 				result := ""
// 				a := fromG + i
// 				if i%100000000 == 0 {
// 					fmt.Println(threadId, a)
// 				}
// 				for range 9 {
// 					b := ((a & 7) ^ 1)
// 					c := a >> b
// 					b ^= 5
// 					a >>= 3
// 					b ^= c
//
// 					result += strconv.Itoa(b & 7)
// 				}
//
// 				if result == "2411751503445530" {
// 					fmt.Println(fromG + i)
// 					panic("done")
// 				}
// 			}
// 		}(threadId)
// 	}
//
// 	time.Sleep(time.Hour * 4)
// }
