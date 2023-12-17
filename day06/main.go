package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

		bar(races)
	case 2:
		tj := strings.Join(times[1:], "")
		dj := strings.Join(dists[1:], "")
		t, _ := strconv.Atoi(tj)
		d, _ := strconv.Atoi(dj)

		bar([]race{{t, d}})
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func foo(races []race) {
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

func bar(races []race) {
	mul := 1

	// hold -> x, time -> t, dist -> d
	// hold * (time - hold) > dist
	// x * (t - x) - d > 0
	// -x^2 + t * x - d > 0, 1 <= x <= t

	for _, r := range races {
		sqrtdelta := math.Sqrt(float64(r.time*r.time - 4*r.dist))
		b := float64(r.time)

		// x1 < x2
		x1f := (b - sqrtdelta) / 2
		x2f := (b + sqrtdelta) / 2

		x1 := int(math.Ceil(x1f))
		x2 := int(math.Floor(x2f))

		if math.Ceil(x1f) == math.Floor(x1f) {
			x1++
		}

		if math.Ceil(x2f) == math.Floor(x2f) {
			x2--
		}

		mul *= x2 - x1 + 1
	}

	fmt.Println(mul)
}
