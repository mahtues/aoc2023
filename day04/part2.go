package main

import (
	"bufio"
	"fmt"
)

func part2(scanner *bufio.Scanner) {
	cards := map[int]int{0: 0}
	n := 0

	for scanner.Scan() {
		cards[n] += 1

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
			for i := n + 1; i < n+1+matches; i++ {
				cards[i] += cards[n]
			}
		}

		n++
	}

	score := 0
	for i := 0; i < n; i++ {
		score += cards[i]
	}

	fmt.Println(score)
}
