package main

import (
	"bufio"
	"fmt"
	"os"
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
	maxSeat := 0

	for _, v := range input {
		rows := v[:7]
		columns := v[7:]

		rowNumber := find(rows, 127, "F", "B")
		columnNumber := find(columns, 7, "L", "R")

		seat := rowNumber*8 + columnNumber
		if seat > maxSeat {
			maxSeat = seat
		}
	}

	return maxSeat
}

func part2(input []string) int {
	list := make([]int, 0)

	for _, v := range input {
		rows := v[:7]
		columns := v[7:]

		rowNumber := find(rows, 127, "F", "B")
		columnNumber := find(columns, 7, "L", "R")

		seat := rowNumber*8 + columnNumber
		list = append(list, seat)
	}

	list = sort(list)

	for i := 1; i < len(list)-1; i++ {
		if list[i] == list[i-1]+2 {
			return list[i] - 1
		}
	}

	return 0
}

func find(input string, max int, left, right string) int {
	min := 0
	for i := 0; i < len(input); i++ {
		if string(input[i]) == left {
			max = (max + min) / 2
			if i == len(input)-1 {
				return max
			}
		} else if string(input[i]) == right {
			min = ((max + min) / 2) + 1
			if i == len(input)-1 {
				return min
			}
		}
	}
	return 0
}

func sort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}

	return input
}
