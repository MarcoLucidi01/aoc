// https://adventofcode.com/2023/day/3
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

	part1, part2 := solve(lines)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, scanner.Err()
}

type pos struct {
	x, y int
}

func solve(lines []string) (int, int) {
	seen := make(map[pos]struct{})
	sum1 := 0
	sum2 := 0
	for i, s := range lines {
		for j, r := range s {
			if r == '.' || unicode.IsDigit(r) {
				continue
			}

			numbers := findAdjacentsNumbers(i, j, lines)
			for start, n := range numbers {
				if _, ok := seen[start]; !ok {
					seen[start] = struct{}{}
					sum1 += n
				}
			}

			if r == '*' && len(numbers) == 2 {
				ratio := 1
				for _, n := range numbers {
					ratio *= n
				}
				sum2 += ratio
			}
		}
	}

	return sum1, sum2
}

func findAdjacentsNumbers(x, y int, lines []string) map[pos]int {
	a := adjacents(x, y)

	numbers := make(map[pos]int)
	for i := 0; i < len(a); i += 2 {
		if x, y := a[i], a[i+1]; unicode.IsDigit(rune(lines[x][y])) {
			y, n := parseNumber(y, lines[x])
			numbers[pos{x, y}] = n
		}
	}

	return numbers
}

func adjacents(x, y int) []int {
	a := []int{-1, 0, +1, 0, 0, -1, 0, +1, -1, -1, -1, +1, +1, -1, +1, +1}

	for i := 0; i < len(a); i += 2 {
		if x+a[i] >= 0 && y+a[i+1] >= 0 {
			a[i] += x
			a[i+1] += y
		} else {
			a = append(a[:i], a[i+2:]...)
		}
	}

	return a
}

func parseNumber(y int, s string) (int, int) {
	start := y
	for i := start - 1; i >= 0 && unicode.IsDigit(rune(s[i])); i-- {
		start--
	}

	end := y
	for i := end + 1; i < len(s) && unicode.IsDigit(rune(s[i])); i++ {
		end++
	}

	n, err := strconv.Atoi(s[start : end+1])
	if err != nil {
		panic(err)
	}

	return start, n
}
