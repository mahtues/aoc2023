package main

import (
	"bufio"
	"fmt"
)

func part2(scanner *bufio.Scanner) {
	var (
		matrix [][]rune
		nums   [][]*int
	)

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	nums = toNums(matrix)

	sum := 0
	I, J := len(matrix), len(matrix[0])
	for i, row := range matrix {
		for j, r := range row {
			if !isSymbol(r) {
				continue
			}

			vals := map[*int]struct{}{}

			for m := max(i-1, 0); m < min(i+2, I); m++ {
				for n := max(j-1, 0); n < min(j+2, J); n++ {
					if nums[m][n] == nil {
						continue
					}
					vals[nums[m][n]] = struct{}{}
				}
			}

			if len(vals) == 2 {
				mul := 1

				for num := range vals {
					mul *= *num
				}

				sum += mul
			}
		}
	}

	fmt.Println(sum)
}

func isStar(r rune) bool {
	return r == '*'
}
