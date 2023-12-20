package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	partArg  = flag.Int("part", 0, "challenge part. possible values: 1, 2")
	cardRank = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)

type player struct {
	cards string
	bid   int
	vec   []int
}

func main() {
	flag.Parse()

	jIsJoker := false

	switch *partArg {
	case 1:
	case 2:
		cardRank['J'] = 1
		jIsJoker = true
	default:
		flag.Usage()
		os.Exit(1)
	}

	players := []player{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(parts[1])
		players = append(players, player{parts[0], bid, eval(parts[0], jIsJoker)})
	}

	sort.SliceStable(players, func(i, j int) bool {
		for k := 0; k < 10; k++ {
			if d := players[i].vec[k] - players[j].vec[k]; d != 0 {
				return d < 0
			}
		}
		return false
	})

	s := uint64(0)
	for i, p := range players {
		fmt.Printf("%s %4d %2v %d\n", p.cards, p.bid, p.vec, len(p.vec))
		s += uint64((i + 1) * p.bid)
	}

	fmt.Println(s)
}

func eval(cards string, jIsJoker bool) []int {
	c := map[rune]int{}

	for _, r := range cards {
		c[r]++
	}

	vec := []int{}

	for k, v := range c {
		if jIsJoker && k == 'J' {
			continue
		}
		vec = append(vec, v)
	}

	// make vec size 5
	for i := len(vec); i < 5; i++ {
		vec = append(vec, 0)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vec)))

	if jIsJoker {
		for i := 0; i < 5 && c['J'] > 0; i++ {
			d := min(5-vec[i], c['J'])
			vec[i] += d
			c['J'] -= d
		}
	}

	car := []int{}
	for _, r := range cards {
		car = append(car, cardRank[r])
	}

	return append(vec, car...)
}
