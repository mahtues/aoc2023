package main

import (
	"bufio"
	"flag"
	"os"
)

var (
	partArg = flag.Int("part", 0, "challenge part. possible values: 1, 2")
	scanner *bufio.Scanner
)

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	switch *partArg {
	case 1:
		part1(scanner)
	case 2:
		part2(scanner)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
