package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0

	for scanner.Scan() {
		rs := []rune(scanner.Text())

		for i := 0; i < len(rs); i++ {
			if isDigit(rs[i]) {
				sum += int(rs[i]-'0') * 10
				break
			}
		}

		for i := len(rs) - 1; i >= 0; i-- {
			if isDigit(rs[i]) {
				sum += int(rs[i] - '0')
				break
			}
		}
	}

	fmt.Println(sum)
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
