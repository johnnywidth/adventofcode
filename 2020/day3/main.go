package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := make([][]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		row := make([]string, 0)
		for i := 0; i < 100; i++ { // TODO: find better number
			for _, v := range text {
				row = append(row, string(v))
			}
		}

		input = append(input, row)
	}

	fmt.Println("Part one = ", part1(input, 3, 1))
	fmt.Println("Part two = ", part2(input))
}

func part1(input [][]string, right, down int) int {
	start := 0
	startDown := 0
	tree := "#"
	count := 0

	for startDown < len(input)-1 {
		start += right
		startDown += down
		position := input[startDown][start]
		if position == tree {
			count++
		}
	}

	return count
}

func part2(input [][]string) int {
	return part1(input, 1, 1) * part1(input, 3, 1) * part1(input, 5, 1) * part1(input, 7, 1) * part1(input, 1, 2)
}
