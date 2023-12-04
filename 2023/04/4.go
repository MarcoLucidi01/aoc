// https://adventofcode.com/2023/day/4
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	cards, err := readCards(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", part1(cards))
	fmt.Println("part2:", part2(cards))
}

func readCards(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)

	var cards []int
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		i := strings.IndexRune(s, ':')
		split := strings.Split(s[i+1:], "|")

		house := make(map[string]struct{})
		for _, n := range strings.Fields(split[1]) {
			house[strings.TrimSpace(n)] = struct{}{}
		}

		m := 0
		for _, n := range strings.Fields(split[0]) {
			if _, ok := house[strings.TrimSpace(n)]; ok {
				m++
			}
		}
		cards = append(cards, m)
	}

	return cards, scanner.Err()
}

func part1(cards []int) int {
	sum := 0
	for _, m := range cards {
		switch {
		case m < 1:
			continue
		case m == 1:
			sum++
		default:
			sum += 1 << (m - 1)
		}
	}
	return sum
}

func part2(cards []int) int {
	copies := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		for j := 0; j < copies[i]+1; j++ {
			for x := i + 1; x <= i+cards[i]; x++ {
				copies[x]++
			}
		}
	}

	sum := 0
	for _, c := range copies {
		sum += c + 1
	}
	return sum
}
