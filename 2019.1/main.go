package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

	fmt.Println("Part one = ", partOne(input))
	fmt.Println("Part two = ", partTwo(input))
}

func partOne(input []string) int {
	var res int

	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		res += fuel(i)
	}

	return int(res)
}

func partTwo(input []string) int {
	var res int

	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		res += totalFuel(i)
	}

	return int(res)
}

func fuel(mass int) int {
	return int(math.Trunc(float64(mass)/3) - 2)
}

func totalFuel(mass int) int {
	f := fuel(mass)
	if f <= 0 {
		return 0
	}

	return int(f) + totalFuel(int(f))
}
