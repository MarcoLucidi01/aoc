// https://adventofcode.com/2021/day/3
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	report, err := readReport(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(report))
	fmt.Printf("part 2: %d\n", partTwo(report))
}

func readReport(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var report []string
	for scanner.Scan() {
		if n := strings.TrimSpace(scanner.Text()); n != "" {
			report = append(report, n)
		}
	}
	return report, scanner.Err()
}

func partOne(report []string) int64 {
	gammaRate := make([]byte, len(report[0]))
	epsilonRate := make([]byte, len(report[0]))
	for i := 0; i < len(report[0]); i++ {
		cnt0, cnt1 := countBitsAt(report, i)
		if cnt0 > cnt1 {
			gammaRate[i] = '0'
			epsilonRate[i] = '1'
		} else {
			gammaRate[i] = '1'
			epsilonRate[i] = '0'
		}
	}
	gamma, _ := strconv.ParseInt(string(gammaRate), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonRate), 2, 64)
	return gamma * epsilon
}

func countBitsAt(report []string, i int) (int, int) {
	cnt0, cnt1 := 0, 0
	for _, n := range report {
		if n[i] == '0' {
			cnt0++
		} else {
			cnt1++
		}
	}
	return cnt0, cnt1
}

func partTwo(report []string) int64 {
	rating := findRating(report, func(cnt0, cnt1 int) rune {
		if cnt0 > cnt1 {
			return '0'
		}
		return '1'
	})
	oxygen, _ := strconv.ParseInt(rating[0], 2, 64)

	rating = findRating(report, func(cnt0, cnt1 int) rune {
		if cnt1 < cnt0 {
			return '1'
		}
		return '0'
	})
	co2, _ := strconv.ParseInt(rating[0], 2, 64)

	return oxygen * co2
}

func findRating(report []string, criteria func(cnt0, cnt1 int) rune) []string {
	rem := make([]string, len(report))
	copy(rem, report)
	for i := 0; i < len(report[0]) && len(rem) > 1; i++ {
		cnt0, cnt1 := countBitsAt(rem, i)
		keep := criteria(cnt0, cnt1)
		for j := 0; j < len(rem); {
			if rune(rem[j][i]) == keep {
				j++
				continue
			}
			rem[j] = rem[len(rem)-1]
			rem = rem[:len(rem)-1]
		}
	}
	return rem
}
