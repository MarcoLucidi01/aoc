// https://adventofcode.com/2024/day/1
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right, err := readLists(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(left, right))
	fmt.Println("part2:", part2(left, right))
}

func readLists(r io.Reader) ([]int, []int, error) {
	scanner := bufio.NewScanner(r)

	var left, right []int
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		n, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, n)

		n, err = strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, err
		}
		right = append(right, n)
	}

	return left, right, scanner.Err()
}

func part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	r := 0
	for i := 0; i < len(left); i++ {
		d := left[i] - right[i]
		if d < 0 {
			d = -d
		}
		r += d
	}
	return r
}

func part2(left, right []int) int {
	freq := make(map[int]int)
	for _, n := range right {
		freq[n]++
	}

	r := 0
	for _, n := range left {
		r += n * freq[n]
	}
	return r
}
