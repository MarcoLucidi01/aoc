// https://adventofcode.com/2022/day/4
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
	part1, part2, err := countOverlap(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func countOverlap(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		split := strings.FieldsFunc(s, func(r rune) bool { return r == ',' || r == '-' })
		if len(split) != 4 {
			return 0, 0, fmt.Errorf("%q: invalid len(split): %d != 4", s, len(split))
		}

		var n [4]int
		for i, sp := range split {
			var err error
			if n[i], err = strconv.Atoi(sp); err != nil {
				return 0, 0, err
			}
		}

		if n[0] <= n[2] && n[1] >= n[3] || n[2] <= n[0] && n[3] >= n[1] {
			sum1++
		}
		if n[2] <= n[1] && n[3] >= n[0] {
			sum2++
		}
	}
	return sum1, sum2, scanner.Err()
}
