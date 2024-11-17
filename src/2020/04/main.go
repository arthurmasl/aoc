package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"aoc/src/internal/utils"
)

var (
	validKeys   = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
)

type Fields map[string]string

type Field struct {
	key, value string
}

func main() {
	lines := utils.GetLines("src/2020/04/input", "\n\n")
	valids := 0

	for _, line := range lines {
		pairs := strings.Split(strings.ReplaceAll(line, "\n", " "), " ")
		fields := make(Fields)
		valid := false

		for _, pair := range pairs {
			key := strings.Split(pair, ":")[0]
			value := strings.Split(pair, ":")[1]
			fields[key] = value
		}

		valid = isValid(fields)

		if valid {
			valids++
		}
	}

	fmt.Println(valids)
}

func isValid(fields Fields) bool {
	keys := []string{}
	for key := range fields {
		keys = append(keys, key)
	}

	for _, validKey := range validKeys {
		if !slices.Contains(keys, validKey) {
			return false
		}
	}

	for key, value := range fields {
		if key == "byr" || key == "iyr" || key == "eyr" {
			year, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

			if key == "byr" {
				if year < 1920 || year > 2002 {
					return false
				}
			}

			if key == "iyr" {
				if year < 2010 || year > 2020 {
					return false
				}
			}

			if key == "eyr" {
				if year < 2020 || year > 2030 {
					return false
				}
			}
		}

		if key == "hgt" {
			mes := value[len(value)-2:]
			if !strings.Contains("cm", mes) && !strings.Contains("in", mes) {
				return false
			}

			height, _ := strconv.Atoi(strings.Split(value, mes)[0])

			if mes == "cm" {
				if height < 150 || height > 193 {
					return false
				}
			}
			if mes == "in" {
				if height < 59 || height > 76 {
					return false
				}
			}
		}

		if key == "hcl" {
			if !isHex(value) {
				return false
			}
		}

		if key == "ecl" {
			if !slices.Contains(validColors, value) {
				return false
			}
		}

		if key == "pid" {
			if utf8.RuneCountInString(value) != 9 {
				return false
			}
			_, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

		}
	}

	return true
}

func isHex(input string) bool {
	re := regexp.MustCompile(`^#[a-fA-F0-9]{6}$`)
	return re.MatchString(input)
}
