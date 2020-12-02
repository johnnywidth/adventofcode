package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		inputNumber, _ := strconv.Atoi(text)
		input = append(input, inputNumber)
	}

	fmt.Println("Part one = ", part1(input))
	fmt.Println("Part twi = ", part2(input))
}

func part1(input []int) int {
	i := 0
	for {
		if i == len(input)-1 {
			break
		}

		for j := i + 1; j < len(input); j++ {
			sum := input[i] + input[j]
			if sum == 2020 {
				return input[i] * input[j]
			}
		}

		i++
	}

	return 0
}

func part2(input []int) int {
	i := 0
	for {
		if i == len(input)-2 {
			break
		}

		for j := i + 1; j < len(input)-1; j++ {
			for k := j + i + 1; k < len(input); k++ {
				sum := input[i] + input[j] + input[k]
				if sum == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}

		i++
	}

	return 0
}
