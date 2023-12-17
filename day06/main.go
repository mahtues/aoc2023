package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	partArg = flag.Int("part", 0, "challenge part. possible values: 1, 2")
	scanner *bufio.Scanner
)

type race struct{ time, dist int }

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	times := strings.Fields(scanner.Text())

	scanner.Scan()
	dists := strings.Fields(scanner.Text())

	races := make([]race, 0, len(times)-1)

	switch *partArg {
	case 1:
		for i := 1; i < len(times); i++ {
			t, _ := strconv.Atoi(times[i])
			d, _ := strconv.Atoi(dists[i])
			races = append(races, race{t, d})
		}

		do(races)
	case 2:
		tj := strings.Join(times[1:], "")
		dj := strings.Join(dists[1:], "")
		t, _ := strconv.Atoi(tj)
		d, _ := strconv.Atoi(dj)

		do([]race{{t, d}})
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func do(races []race) {
	mul := 1

	for _, r := range races {
		count := 0

		for hold := 1; hold < r.time; hold++ {
			if dist(hold, r.time) > r.dist {
				count++
			}
		}

		mul *= count
	}

	fmt.Println(mul)
}

func dist(hold int, time int) int {
	return hold * (time - hold)
}
