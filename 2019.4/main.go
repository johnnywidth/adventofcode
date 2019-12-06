package main

import (
	"fmt"
)

func main() {
	input := []int{128392, 643281}
	fmt.Println("Part one = ", partOne(input))
}

func partOne(input []int) int {
	var res int

	for i := input[0]; i <= input[1]; i++ {
		if meetsCriteria(i) {
			res++
		}
	}

	return res
}

func meetsCriteria(pass int) bool {
	originalPass := pass
	digs := make([]int, 0)

	var hasDouble bool
	increaseCount := 1
	var previousDig = -1

	for {
		dig := originalPass

		if originalPass > 10 {
			dig = originalPass % 10
		}

		if previousDig != -1 {
			if dig == previousDig {
				hasDouble = true
			}

			if dig <= previousDig {
				increaseCount++
			}
		}

		previousDig = dig
		digs = append(digs, dig)

		if originalPass > 10 {
			originalPass = (originalPass - dig) / 10
		} else {
			break
		}
	}

	return hasDouble && increaseCount == 6 && len(digs) == 6
}
