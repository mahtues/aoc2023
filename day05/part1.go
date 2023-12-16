package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type state int

const (
	seed = iota
	seed2soil
	soil2fert
	fert2water
	water2light
	light2temp
	temp2humid
	humid2loc
)

var stateString = map[state]string{
	seed:        "seed",
	seed2soil:   "seed2soil",
	soil2fert:   "soil2fert",
	fert2water:  "fert2water",
	water2light: "water2light",
	light2temp:  "light2temp",
	temp2humid:  "temp2humid",
	humid2loc:   "humid2loc",
}

type entry struct{ start, target, length int }

func part1(scanner *bufio.Scanner) {
	state := seed

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")[1:]
	seeds := make([]int, 0, len(parts))
	for _, s := range parts {
		x, _ := strconv.Atoi(s)
		seeds = append(seeds, x)
	}

	mappings := []entry{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanner.Scan() // skip header
			if state > seed {
				next := make([]int, 0, len(seeds))
				for _, s := range seeds {
					next = append(next, forward(s, mappings))
				}
				seeds = next
			}
			state++
			clear(mappings)
			continue
		}

		parts := strings.Split(scanner.Text(), " ")

		target, _ := strconv.Atoi(parts[0])
		source, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])

		mappings = append(mappings, entry{source, target, length})
	}

	next := make([]int, 0, len(seeds))
	for _, s := range seeds {
		next = append(next, forward(s, mappings))
	}

	x := next[0]
	for _, y := range next[1:] {
		x = min(x, y)
	}

	fmt.Println(x)
}

func forward(val int, mappings []entry) int {
	for _, m := range mappings {
		if m.start <= val && val < m.start+m.length {
			return val - m.start + m.target
		}
	}

	return val
}

func splitch(s string, seps []rune) <-chan string {
	sch := make(chan string)

	sepmap := map[rune]struct{}{}

	for _, r := range seps {
		sepmap[r] = struct{}{}
	}

	go func() {
		rs := make([]rune, 0, len(s))

		for _, r := range s {
			if _, ok := sepmap[r]; ok {
				if len(rs) > 0 {
					sch <- string(rs)
					clear(rs)
				}
			} else {
				rs = append(rs, r)
			}
		}

		if len(rs) > 0 {
			sch <- string(rs)
		}

		close(sch)
	}()

	return sch
}
