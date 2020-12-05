package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var requiredFileds = map[string]struct{}{
	"byr": {},
	"iyr": {},
	"eyr": {},
	"hgt": {},
	"hcl": {},
	"ecl": {},
	"pid": {},
}

func main() {
	input := make([][]string, 0)
	row := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			input = append(input, row)
			row = make([]string, 0)
			continue
		}

		for _, v := range strings.Split(text, " ") {
			row = append(row, string(v))
		}
	}

	fmt.Println("Part one = ", part1(input))
	fmt.Println("Part two = ", part2(input))
}

func part1(input [][]string) int {
	count := 0

	for _, row := range input {
		preCount := 0
		for _, v := range row {
			keyValue := strings.Split(v, ":")
			if _, ok := requiredFileds[keyValue[0]]; ok {
				preCount++
			}
		}
		if preCount == len(requiredFileds) {
			count++
		}
	}

	return count
}

func part2(input [][]string) int {
	count := 0

	for _, row := range input {
		preCount := 0
		var notValid bool
		for _, v := range row {
			keyValue := strings.Split(v, ":")
			if _, ok := requiredFileds[keyValue[0]]; ok {
				preCount++
			}

			if !validation(keyValue[0], keyValue[1]) {
				fmt.Println("not valid: ", keyValue[0], keyValue[1])
				notValid = true
			}
		}

		if notValid || preCount != len(requiredFileds) {
			continue
		}
		count++
	}

	return count
}

func validation(key, value string) bool {
	switch key {
	case "byr":
		valueNumber, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if len(value) < 4 || len(value) > 4 ||
			valueNumber < 1920 || valueNumber > 2002 {
			return false
		}
	case "iyr":
		valueNumber, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if len(value) < 4 || len(value) > 4 ||
			valueNumber < 2010 || valueNumber > 2020 {
			return false
		}
	case "eyr":
		valueNumber, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if len(value) < 4 || len(value) > 4 ||
			valueNumber < 2020 || valueNumber > 2030 {
			return false
		}
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			newValue := strings.TrimSuffix(value, "cm")
			valueNumber, err := strconv.Atoi(newValue)
			if err != nil {
				panic(err)
			}
			if valueNumber < 150 || valueNumber > 193 {
				return false
			}
		} else if strings.HasSuffix(value, "in") {
			newValue := strings.TrimSuffix(value, "in")
			valueNumber, err := strconv.Atoi(newValue)
			if err != nil {
				panic(err)
			}
			if valueNumber < 59 || valueNumber > 76 {
				return false
			}
		} else {
			return false
		}
	case "hcl":
		m, err := regexp.Match("^#(?:[0-9a-fA-F]{6})$", []byte(value))
		if err != nil {
			panic(err)
		}
		return m
	case "ecl":
		cases := map[string]struct{}{
			"amb": {},
			"blu": {},
			"brn": {},
			"gry": {},
			"grn": {},
			"hzl": {},
			"oth": {},
		}
		_, ok := cases[value]
		return ok
	case "pid":
		m, err := regexp.Match(`^\d{9}$`, []byte(value))
		if err != nil {
			panic(err)
		}
		return m
	case "cid":
		return true
	default:
		return false
	}

	return true
}
