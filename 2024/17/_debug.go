package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

const (
	from = 36000000000000
	to   = 290000000000000
)

const (
	exampleInput  = 23999685
	exampleOutput = "503576154"
	target        = "2411751503445530"
)

var targetArr = []int{2, 4, 1, 1, 7, 5, 1, 5, 0, 3, 4, 4, 5, 5, 3, 0}

var (
	color = "\033[41m"
	reset = "\033[0m"
)

// wow
// input := 164525255782333
// wow2 164525658435517
func main() {
	//                 x
	//       164525658435517
	input := 164525658435517
	//       164525658435517
	output := program(input)

	// change - 1645256 / 58435517

	// fmt.Println("input ", input)
	fmt.Println("output", output[0:10], output[10:11], output[11:])
	fmt.Println("target", target[0:10], target[10:11], target[11:])
	// fmt.Println("target", target)
	fmt.Println()

	utils.Assert(strings.Contains(program(input), "2411751503"))
	utils.Assert(strings.Contains(program(input), "45530"))

	// utils.Assert(program(exampleInput) == exampleOutput)
	// utils.Assert(toInt(program(reverseProgram(targetArr))) == target)

	// bruteforce()
	// letsgo()
	// debug()
}

func program(a int) string {
	var result string

	for a > 0 {
		b := a&7 ^ 1
		c := a >> b
		a >>= 3
		b ^= 5 ^ c

		output := b & 7
		result += strconv.Itoa(output)
	}

	return result
}

func debug() {
	//         164525658435261
	initial := 164625658435261
	step := 1

	a := initial

	reader := bufio.NewReader(os.Stdin)
	printDebug(a, step)

	// return
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			printDebug(a, step)
			a += step
			continue
		}

		if input == "r" {
			step = -step
			continue
		}
		if input == "i" {
			a = initial
			continue
		}
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		step = inputInt
	}
}

func printDebug(a, step int) {
	result := program(a)
	if result == target {
		fmt.Println("===", a)
		panic("found")
	}

	clearConsole()
	fmt.Println("step  ", step)
	fmt.Println("number", a)
	fmt.Println("target", target)
	// fmt.Println("output", result)
	fmt.Print("output ")
	for i, c := range result {
		if i >= len(target) {
			fmt.Print(string(c))
			continue
		}
		if target[i] == byte(c) {
			fmt.Printf("%v%v%v", color, string(c), reset)
			continue
		}
		fmt.Print(string(c))
	}
	fmt.Println()
	fmt.Print("input  ")
}

func letsgo() {
	// change - 1645256 / 58435517
	left := "164"
	// right := "58435517"

	i := 100000000000
	for {
		i++
		numberStr := left + strconv.Itoa(i)
		number, _ := strconv.Atoi(numberStr)

		result := program(number)
		// fmt.Println(number)
		// fmt.Println(target)
		// fmt.Println(result)
		// fmt.Println()

		if result == target {
			fmt.Println(number)
			break
		}
		if strings.Contains(result, "2411751503") {
			fmt.Println(result, number)
		}
	}
}

func bruteforce() {
	//   164525658435517
	a := 165525658435517
	for {
		a += 1
		result := program(a)
		if result == target {
			fmt.Println("=== found!!", a)
			break
		}

		//                                      x
		//                           2411751503445530
		if strings.Contains(result, "241175150") {
			fmt.Println(result, a)
		}
	}

	// time.Sleep(time.Hour)
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func toInt(ints []int) int {
	str := ""
	for _, n := range ints {
		str += strconv.Itoa(n)
	}
	number, _ := strconv.Atoi(str)
	return number
}
