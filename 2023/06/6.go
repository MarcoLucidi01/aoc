// https://adventofcode.com/2023/day/6
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
	times, time, distances, dist, err := readPastRaces(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", part1(times, distances))
	fmt.Println("part2:", part2(time, dist))
}

func readPastRaces(r io.Reader) ([]int, int, []int, int, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	times, time, err := parseNumbers(scanner.Text())
	if err != nil {
		return nil, 0, nil, 0, err
	}

	scanner.Scan()
	distances, dist, err := parseNumbers(scanner.Text())
	if err != nil {
		return nil, 0, nil, 0, err
	}

	return times, time, distances, dist, nil
}

func parseNumbers(s string) ([]int, int, error) {
	s = strings.TrimSpace(s)
	i := strings.IndexRune(s, ':')
	fields := strings.Fields(s[i+1:])

	var nums []int
	for _, f := range fields {
		n, err := strconv.Atoi(strings.TrimSpace(f))
		if err != nil {
			return nil, 0, err
		}
		nums = append(nums, n)
	}

	n, err := strconv.Atoi(strings.Join(fields, ""))
	if err != nil {
		return nil, 0, err
	}

	return nums, n, nil
}

func part1(times []int, distances []int) int {
	ans := 1
	for i, t := range times {
		n := 0
		for j := 0; j < t; j++ {
			if (t-j)*j > distances[i] {
				n++
			}
		}
		ans *= n
	}
	return ans
}

func part2(time int, dist int) int {
	n := 0
	for j := 0; j < time; j++ {
		if (time-j)*j > dist {
			n++
		}
	}
	return n
}
