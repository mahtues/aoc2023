package main

import (
	"bufio"
	"fmt"
)

type state int

const (
	nothing state = iota
	parsingNumber
)

func part1(scanner *bufio.Scanner) {
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

			for m := max(i-1, 0); m < min(i+2, I); m++ {
				for n := max(j-1, 0); n < min(j+2, J); n++ {
					if nums[m][n] == nil {
						continue
					}

					sum += *nums[m][n]
					*nums[m][n] = 0
				}
			}
		}
	}

	fmt.Println(sum)
}

func toDigit(r rune) (int, bool) {
	if '0' <= r && r <= '9' {
		return int(r - '0'), true
	}

	return 0, false
}

func isSymbol(r rune) bool {
	if _, b := toDigit(r); b {
		return false
	}

	if r == '.' {
		return false
	}

	return true
}

func toNums(matr [][]rune) [][]*int {
	var (
		state state = nothing
		num   *int  = nil
		nums  [][]*int
	)

	nums = make([][]*int, len(matr))
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]*int, len(matr[0]))
	}

	for i, row := range matr {
		for j, r := range row {
			switch state {
			case nothing:
				if x, ok := toDigit(r); ok {
					num = new(int)
					*num = x
					nums[i][j] = num

					state = parsingNumber
				} else {
					state = nothing
				}
			case parsingNumber:
				if x, ok := toDigit(r); ok {
					*num = *num*10 + x
					nums[i][j] = num

					state = parsingNumber
				} else {
					state = nothing
				}
			}
		}
	}

	return nums
}
