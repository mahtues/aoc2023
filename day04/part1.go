package main

import (
	"bufio"
	"fmt"
)

type state int

type stage int

const (
	findingColon state = iota
	findingNumber
	parsingNumber
)

const (
	winners stage = iota
	possesses
)

func part1(scanner *bufio.Scanner) {
	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		var (
			state   state            = findingColon
			stage   stage            = winners
			winning map[int]struct{} = map[int]struct{}{}
			num     int              = 0
			matches int              = 0
		)

		for _, r := range line + " " {
			hold := true
			for hold {
				hold = false

				switch state {
				case findingColon:
					if r == ':' {
						state = findingNumber
						continue
					}

				case findingNumber:
					if _, ok := toDigit(r); ok {
						num = 0
						state = parsingNumber
						hold = true
						continue
					}

					if r == '|' {
						stage = possesses
					}

				case parsingNumber:
					if x, ok := toDigit(r); ok {
						num = num*10 + x
						continue
					}

					hold = true
					state = findingNumber

					if stage == winners {
						winning[num] = struct{}{}
						continue
					}

					if _, ok := winning[num]; ok {
						matches += 1
					}
				}
			}
		}

		if matches > 0 {
			score += 1 << (matches - 1)
		}
	}

	fmt.Println(score)
}

func toDigit(r rune) (int, bool) {
	if '0' <= r && r <= '9' {
		return int(r - '0'), true
	}

	return 0, false
}
