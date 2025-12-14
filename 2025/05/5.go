// https://adventofcode.com/2025/day/5
package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	ranges, available, err := readIngredients(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d\n", part1(ranges, available))
	fmt.Printf("part2: %d\n", part2(ranges))
}

func readIngredients(r io.Reader) ([][2]int, []int, error) {
	scanner := bufio.NewScanner(r)

	var ranges [][2]int
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			break
		}

		split := strings.Split(s, "-")
		if len(split) != 2 {
			return nil, nil, fmt.Errorf("invalid range %q", s)
		}

		a, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, nil, err
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, nil, err
		}

		ranges = append(ranges, [2]int{a, b})
	}

	var available []int
	for scanner.Scan() {
		id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, nil, err
		}
		available = append(available, id)
	}

	return ranges, available, scanner.Err()
}

func part1(ranges [][2]int, available []int) int {
	n := 0
loop:
	for _, a := range available {
		for _, r := range ranges {
			if a >= r[0] && a <= r[1] {
				n++
				continue loop
			}
		}
	}

	return n
}

func part2(ranges [][2]int) int {
	slices.SortFunc(ranges, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })

	n := 0
	var newRanges [][2]int
loop:
	for _, r := range ranges {
		a, b := r[0], r[1]
		for _, nr := range newRanges {
			if a >= nr[0] && a <= nr[1] {
				a = nr[1] + 1
				if a > b {
					continue loop
				}
			}
		}

		n += b - a + 1
		newRanges = append(newRanges, [2]int{a, b})
	}

	return n
}
