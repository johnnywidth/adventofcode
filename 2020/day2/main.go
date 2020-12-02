package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		input = append(input, text)
	}

	fmt.Println("Part one = ", part1(input))
	fmt.Println("Part two = ", part2(input))
}

func part1(input []string) int {
	var count int
	for _, v := range input {
		raw := strings.Split(v, " ")
		policy := strings.Split(raw[0], "-")

		min, _ := strconv.Atoi(policy[0])
		max, _ := strconv.Atoi(policy[1])
		letter := raw[1][0]
		password := raw[2]

		var i int
		for _, pl := range password {
			if byte(pl) == letter {
				i++
			}
		}

		if i >= min && i <= max {
			count++
		}
	}

	return count
}

func part2(input []string) int {
	var count int
	for _, v := range input {
		raw := strings.Split(v, " ")
		policy := strings.Split(raw[0], "-")

		positionOne, _ := strconv.Atoi(policy[0])
		positionTwo, _ := strconv.Atoi(policy[1])
		letter := raw[1][0]
		password := raw[2]

		if (password[positionOne-1] == letter &&
			password[positionTwo-1] != letter) ||
			(password[positionOne-1] != letter &&
				password[positionTwo-1] == letter) {
			count++
		}
	}

	return count
}
