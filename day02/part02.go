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

	sum := 0

	for scanner.Scan() {
		var (
			state  = findingGameId
			gameId = 0
			cubes  = 0
			maxes  = map[rune]int{'r': 0, 'g': 0, 'b': 0}
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
						maxes[r] = cubes
					}
					cubes = 0
					state = findingCubes
				}
			}
		}

		sum += maxes['r'] * maxes['g'] * maxes['b']
	}

	fmt.Println(sum)
}

func toDigit(r rune) (int, bool) {
	if '0' <= r && r <= '9' {
		return int(r - '0'), true
	}

	return 0, false
}
