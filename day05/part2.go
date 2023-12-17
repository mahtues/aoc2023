package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type pair struct{ start, length int }

func part2(scanner *bufio.Scanner) {
	state := seed

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")[1:]
	seeds := []pair{}
	for i := 0; i < len(parts); i = i + 2 {
		s, _ := strconv.Atoi(parts[i])
		l, _ := strconv.Atoi(parts[i+1])
		seeds = append(seeds, pair{s, l})
	}

	mappings := []entry{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if state > seed {
				next := []pair{}
				for _, p := range seeds {
					next = append(next, forward2(p, mappings)...)
				}
				// fmt.Printf("[%11s] %+v -> %+v\n", stateString[state], seeds, next)
				seeds = next
			}
			state++
			mappings = []entry{}
			scanner.Scan() // skip header
			continue
		}

		parts := strings.Split(scanner.Text(), " ")

		target, _ := strconv.Atoi(parts[0])
		source, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])

		mappings = append(mappings, entry{source, target, length})
	}

	next := []pair{}
	for _, s := range seeds {
		next = append(next, forward2(s, mappings)...)
	}
	// fmt.Printf("[%10s] %+v -> %+v\n", stateString[state(stt)], seeds, next)

	x := next[0]
	for _, y := range next[1:] {
		if y.start < x.start {
			x = y
		}
	}

	fmt.Println(x.start)
}

func forward2(v pair, mappings []entry) []pair {
	result := []pair{}

	curr := []pair{v}
	for _, b := range mappings {
		next := []pair{}

		for _, a := range curr {
			// b fully overlaps a
			if b.start <= a.start && b.start+b.length >= a.start+a.length {
				result = append(result, pair{a.start - b.start + b.target, a.length})
				// fmt.Println("+", "a:", a, "b:", b, "overlap:", result[len(result)-1])
				break
			}

			// b overlaps a left side
			if b.start <= a.start && b.start+b.length > a.start {
				result = append(result, pair{a.start - b.start + b.target, b.start + b.length - a.start})
				next = append(next, pair{b.start + b.length, a.start + a.length - b.start - b.length})
				// fmt.Println("<", "a:", a, "b:", b, "overlap:", result[len(result)-1], "rest:", next[len(next)-1])
				break
			}

			// b overlaps a right side
			if a.start <= b.start && a.start+a.length > b.start {
				result = append(result, pair{b.target, a.start + a.length - b.start})
				next = append(next, pair{a.start, b.start - a.start})
				// fmt.Println(">", "a:", a, "b:", b, "overlap:", result[len(result)-1], "rest:", next[len(next)-1])
				break
			}

			// b inside a
			if a.start < b.start && a.start+a.length > b.start+b.length {
				result = append(result, pair{b.target, b.length})
				next = append(next, pair{a.start, b.start - a.start}, pair{b.start + b.length, a.start + a.length - b.start - b.length})
				// fmt.Println("-", "a:", a, "b:", b, "overlap:", result[len(result)-1], "rest:", next[len(next)-2:])
				break
			}

			// fmt.Println(" ", "a:", a, "b:", b)
			next = append(next, a)
		}

		curr = next
		if len(curr) == 0 {
			break
		}
	}

	result = append(result, curr...)
	// fmt.Println("result", result)

	return result
}
