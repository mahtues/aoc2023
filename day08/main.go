package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var partArg = flag.Int("part", 0, "challenge part. possible values: 1, 2")

type side struct{ left, right string }

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	switch *partArg {
	case 1:
		part1(scanner)
	case 2:
		part2smart(scanner)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func part1(scanner *bufio.Scanner) {
	scanner.Scan()
	path := scanner.Text()

	scanner.Scan() // skip blank line

	r, _ := regexp.Compile(`(?P<curr>\w+) = \((?P<left>\w+), (?P<right>\w+)\)`)

	network := map[string]side{}

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindStringSubmatch(line)

		nodes := map[string]string{}

		for i, name := range r.SubexpNames() {
			if name != "" {
				nodes[name] = matches[i]
			}
		}

		network[nodes["curr"]] = side{nodes["left"], nodes["right"]}
	}

	curr := "AAA"
	count := 0
	for {
		for _, s := range path {
			if s == 'L' {
				curr = network[curr].left
			} else {
				curr = network[curr].right
			}
			count++

			if curr == "ZZZ" {
				break
			}
		}

		if curr == "ZZZ" {
			break
		}
	}

	fmt.Println(count)
}

func part2smart(scanner *bufio.Scanner) {
	scanner.Scan()
	path := scanner.Text()

	scanner.Scan() // skip blank line

	r, _ := regexp.Compile(`(?P<curr>\w+) = \((?P<left>\w+), (?P<right>\w+)\)`)

	var (
		network = map[string]side{}
		currs   = []string{}
	)

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindStringSubmatch(line)

		nodes := map[string]string{}

		for i, name := range r.SubexpNames() {
			if name != "" {
				nodes[name] = matches[i]
			}
		}

		node := nodes["curr"]
		network[node] = side{nodes["left"], nodes["right"]}

		rs := []rune(node)
		if rs[len(rs)-1] == 'A' {
			currs = append(currs, node)
		}
	}

	var (
		step  = 0
		stop  = false
		steps = []uint64{}
	)
	for {
		for _, s := range path {
			step++

			nexts := []string{}

			for _, curr := range currs {
				var next string
				if s == 'L' {
					next = network[curr].left
				} else {
					next = network[curr].right
				}

				rs := []rune(next)
				if rs[len(rs)-1] == 'Z' {
					steps = append(steps, uint64(step))
				} else {
					nexts = append(nexts, next)
				}
			}

			stop = len(nexts) == 0

			if stop {
				break
			}

			currs = nexts
		}

		if stop {
			break
		}
	}

	for len(steps) > 1 {
		nexts := []uint64{}

		if len(steps)%2 == 1 {
			nexts = append(nexts, steps[len(steps)-1])
		}

		for i := 0; i < len(steps)-1; i += 2 {
			nexts = append(nexts, lcm(steps[i], steps[i+1]))
		}

		steps = nexts
	}

	fmt.Println(steps[0])
}

func part2dumb(scanner *bufio.Scanner) {
	scanner.Scan()
	path := scanner.Text()

	scanner.Scan() // skip blank line

	r, _ := regexp.Compile(`(?P<curr>\w+) = \((?P<left>\w+), (?P<right>\w+)\)`)

	network := map[string]side{}
	currs := map[string]struct{}{}

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindStringSubmatch(line)

		nodes := map[string]string{}

		for i, name := range r.SubexpNames() {
			if name != "" {
				nodes[name] = matches[i]
			}
		}

		network[nodes["curr"]] = side{nodes["left"], nodes["right"]}

		rs := []rune(nodes["curr"])
		if rs[len(rs)-1] == 'A' {
			currs[nodes["curr"]] = struct{}{}
		}
	}

	count := 0
	stop := false
	for {
		for _, s := range path {
			nexts := map[string]struct{}{}
			zs := 0

			for curr := range currs {
				var next string

				if s == 'L' {
					next = network[curr].left
				} else {
					next = network[curr].right
				}

				nexts[next] = struct{}{}

				rs := []rune(next)
				if rs[len(rs)-1] == 'Z' {
					zs++
				}
			}

			count++

			if zs > 1 {
				fmt.Printf("%10d %v %*d\n", count, nexts, zs, zs)
			}

			stop = len(currs) == zs

			if stop {
				break
			}

			currs = nexts
		}

		if stop {
			break
		}
	}

	fmt.Println(count)
}

func gcd(a uint64, b uint64) uint64 {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a uint64, b uint64) uint64 {
	if b > a {
		a, b = b, a
	}

	return a / gcd(a, b) * b
}
