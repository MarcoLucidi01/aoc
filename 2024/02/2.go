// https://adventofcode.com/2024/day/2
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
	reports, err := readReports(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(reports))
	fmt.Println("part2:", part2(reports))
}

func readReports(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)

	var reports [][]int
	for scanner.Scan() {
		var report []int
		for _, f := range strings.Fields(scanner.Text()) {
			n, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}

	return reports, scanner.Err()
}

func part1(reports [][]int) int {
	n := 0
	for _, report := range reports {
		if isSafe(report) {
			n++
		}
	}
	return n
}

func part2(reports [][]int) int {
	n := 0
	for _, report := range reports {
		for i := 0; i < len(report); i++ {
			if isSafe(append(append([]int(nil), report[:i]...), report[i+1:]...)) {
				n++
				break
			}
		}
	}
	return n
}

func isSafe(report []int) bool {
	sign := '+'
	for i := 1; i < len(report); i++ {
		s := '+'
		d := report[i-1] - report[i]
		if d < 0 {
			s = '-'
			d = -d
		}

		if i == 1 {
			sign = s
		}

		if s != sign || d < 1 || d > 3 {
			return false
		}
	}
	return true
}
