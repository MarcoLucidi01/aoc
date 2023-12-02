// https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines, err := readLines(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	p1, err1 := part1(lines)
	fmt.Println("part1:", p1, "err:", err1)

	p2, err2 := part2(lines)
	fmt.Println("part2:", p2, "err:", err2)
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, scanner.Err()
}

func part1(lines []string) (int, error) {
	sum := 0
	for _, s := range lines {
		n, err := findNumber(s)
		if err != nil {
			return 0, err
		}
		sum += n
	}

	return sum, nil
}

func findNumber(s string) (int, error) {
	a := strings.IndexFunc(s, unicode.IsDigit)
	if a < 0 {
		return 0, fmt.Errorf("%s: no first digit", s)
	}

	b := strings.LastIndexFunc(s, unicode.IsDigit)
	if b < 0 {
		return 0, fmt.Errorf("%s: no last digit", s)
	}

	return strconv.Atoi(string(s[a]) + string(s[b]))
}

func part2(lines []string) (int, error) {
	sum := 0
	for _, s := range lines {
		s = strings.ReplaceAll(s, "one", "o1e")
		s = strings.ReplaceAll(s, "two", "t2o")
		s = strings.ReplaceAll(s, "three", "t3e")
		s = strings.ReplaceAll(s, "four", "f4r")
		s = strings.ReplaceAll(s, "five", "f5e")
		s = strings.ReplaceAll(s, "six", "s6x")
		s = strings.ReplaceAll(s, "seven", "s7n")
		s = strings.ReplaceAll(s, "eight", "e8t")
		s = strings.ReplaceAll(s, "nine", "n9e")

		n, err := findNumber(s)
		if err != nil {
			return 0, err
		}
		sum += n
	}

	return sum, nil
}
