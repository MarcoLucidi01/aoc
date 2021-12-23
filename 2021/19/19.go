// https://adventofcode.com/2021/day/19
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readReports()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(reports))
	fmt.Printf("part 2: %d\n", partTwo(reports))
}

func readReports() ([][][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var reports [][][]int
	var report [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "--- scanner ") {
			if len(report) > 0 {
				reports = append(reports, report)
			}
			report = nil
			continue
		}
		var numbers []int
		for _, s := range strings.Split(line, ",") {
			if n, err := strconv.Atoi(s); err == nil {
				numbers = append(numbers, n)
			}
		}
		report = append(report, numbers)
	}
	if len(report) > 0 {
		reports = append(reports, report)
	}
	return reports, scanner.Err()
}

func partOne(reports [][][]int) int {
	// TODO
	return 0
}

func partTwo(reports [][][]int) int {
	// TODO
	return 0
}
