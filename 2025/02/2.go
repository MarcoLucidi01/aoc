// https://adventofcode.com/2025/day/2
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, err := readRanges(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d\n", part1(ranges))
	fmt.Printf("part2: %d\n", part2(ranges))
}

func readRanges(r io.Reader) ([][2]int, error) {
	var ranges [][2]int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		split := strings.FieldsFunc(s, func(r rune) bool { return r == ',' || r == '-' })
		for i := 0; i < len(split); i += 2 {
			var err error
			var r [2]int
			if r[0], err = strconv.Atoi(split[i]); err != nil {
				return nil, err
			}
			if r[1], err = strconv.Atoi(split[i+1]); err != nil {
				return nil, err
			}
			ranges = append(ranges, r)
		}
	}

	return ranges, scanner.Err()
}

func part1(ranges [][2]int) int {
	sum := 0
	for _, r := range ranges {
		for n := r[0]; n <= r[1]; n++ {
			s := strconv.Itoa(n)
			if len(s)%2 != 0 {
				continue
			}
			if s[:len(s)/2] == s[len(s)/2:] {
				sum += n
			}
		}
	}
	return sum
}

func part2(ranges [][2]int) int {
	sum := 0
	for _, r := range ranges {
		for n := r[0]; n <= r[1]; n++ {
			s := strconv.Itoa(n)
			for i := 1; i < len(s); i++ {
				prefix := s[:i]
				rest := s[i:]
				patterns := 1
				for rest != "" && strings.HasPrefix(rest, prefix) {
					rest = rest[len(prefix):]
					patterns++
				}
				if rest == "" && patterns > 1 {
					sum += n
					break
				}
			}
		}
	}
	return sum
}
