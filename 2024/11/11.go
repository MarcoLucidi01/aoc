// https://adventofcode.com/2024/day/11
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
	stones, err := readStones(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", solve(stones, 25))
	fmt.Println("part2:", solve(stones, 75))
}

func readStones(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var stones []int
	for i := 0; scanner.Scan(); i++ {
		for _, s := range strings.Fields(strings.TrimSpace(scanner.Text())) {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			stones = append(stones, int(n))
		}
	}
	return stones, scanner.Err()
}

type entry struct {
	n, step int
}

func solve(stones []int, max int) int {
	cache := make(map[entry]int)

	n := 0
	for _, s := range stones {
		n += 1 + blink(s, 0, max, cache)
	}
	return n
}

func blink(stone, step, max int, cache map[entry]int) int {
	if step == max {
		return 0
	}

	if n, ok := cache[entry{stone, step}]; ok {
		return n
	}

	if stone == 0 {
		n := blink(1, step+1, max, cache)
		cache[entry{stone, step}] = n
		return n
	}

	if s := strconv.Itoa(stone); len(s)%2 == 0 {
		a, _ := strconv.Atoi(s[:len(s)/2])
		b, _ := strconv.Atoi(s[len(s)/2:])
		n := 1 + blink(a, step+1, max, cache) + blink(b, step+1, max, cache)
		cache[entry{stone, step}] = n
		return n
	}

	n := blink(stone*2024, step+1, max, cache)
	cache[entry{stone, step}] = n
	return n
}
