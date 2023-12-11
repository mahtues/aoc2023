package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0

	for scanner.Scan() {
		s := scanner.Text()

		digits := []int{}

		for i := range s {
			if x, err := toDigit(s, i); err == nil {
				digits = append(digits, x)
			}
		}

		sum += digits[0]*10 + digits[len(digits)-1]
	}

	fmt.Println(sum)
}

func toDigit(s string, i int) (int, error) {
	if '0' <= s[i] && s[i] <= '9' {
		return int(s[i] - '0'), nil
	}

	ds := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for x := 1; x < 10; x++ {
		d := ds[x]

		if i+len(d) > len(s) {
			continue
		}

		s := s[i : i+len(d)]

		if s == d {
			return x, nil
		}
	}

	return 0, errors.New("not a digit")
}
