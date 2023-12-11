// https://adventofcode.com/2023/day/9
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
	values, err := readValues(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	p1, p2 := solve(values)
	fmt.Println("part1:", p1)
	fmt.Println("part2:", p2)
}

func readValues(r io.Reader) ([][]int, error) {
	var values [][]int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var row []int
		for _, s := range strings.Fields(scanner.Text()) {
			n, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}
		values = append(values, row)
	}
	return values, scanner.Err()
}

func solve(values [][]int) (int, int) {
	sum1, sum2 := 0, 0

	for _, row := range values {
		var diffs [][4]int
	loop:
		diffs = append(diffs, [4]int{0, row[0], row[len(row)-1], 0})

		var diff []int
		for i := 1; i < len(row); i++ {
			diff = append(diff, row[i]-row[i-1])
		}
		row = diff

		for _, n := range diff {
			if n != 0 {
				goto loop
			}
		}
		diffs = append(diffs, [4]int{})

		for i := len(diffs) - 2; i >= 0; i-- {
			diffs[i][3] = diffs[i][2] + diffs[i+1][3]
			diffs[i][0] = diffs[i][1] - diffs[i+1][0]
		}
		sum1 += diffs[0][3]
		sum2 += diffs[0][0]
	}

	return sum1, sum2
}
