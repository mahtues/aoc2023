package main

import (
	"bufio"
	"fmt"
	"os"
)

type State int

const (
	findingGameId = iota
	parsingGameId
	findingCubes
	parsingNumber
	findingColor
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		maxes = map[rune]int{'r': 12, 'g': 13, 'b': 14}
		sum   = 0
	)

	for scanner.Scan() {
		var (
			state      = findingGameId
			gameId     = 0
			cubes      = 0
			impossible = false
		)

		for _, r := range scanner.Text() {
			switch state {
			case findingGameId:
				if x, ok := toDigit(r); ok {
					gameId = gameId*10 + x
					state = parsingGameId
				}
			case parsingGameId:
				if x, ok := toDigit(r); ok {
					gameId = gameId*10 + x
				} else {
					state = findingCubes
				}
			case findingCubes:
				if x, ok := toDigit(r); ok {
					cubes = cubes*10 + x
					state = parsingNumber
				}
			case parsingNumber:
				if x, ok := toDigit(r); ok {
					cubes = cubes*10 + x
				} else {
					state = findingColor
				}
			case findingColor:
				if maxCubes, ok := maxes[r]; ok {
					if cubes > maxCubes {
						impossible = true
						break
					} else {
						state = findingCubes
						cubes = 0
					}
				}
			}
		}

		if !impossible {
			sum += gameId
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
