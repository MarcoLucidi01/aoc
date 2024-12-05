// https://adventofcode.com/2024/day/5
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	ordering, updates, err := readPrintQueue(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(ordering, updates))
	fmt.Println("part2:", part2(ordering, updates))
}

func readPrintQueue(r io.Reader) (map[int][]int, [][]int, error) {
	scanner := bufio.NewScanner(r)

	ordering := make(map[int][]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		s := strings.Split(scanner.Text(), "|")
		a, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, nil, err
		}
		b, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, nil, err
		}

		ordering[a] = append(ordering[a], b)
	}

	var updates [][]int
	for scanner.Scan() {
		var pages []int
		s := strings.Split(scanner.Text(), ",")
		for _, v := range s {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, nil, err
			}
			pages = append(pages, n)
		}
		updates = append(updates, pages)
	}

	return ordering, updates, nil
}

func part1(ordering map[int][]int, updates [][]int) int {
	sum := 0
	for _, pages := range updates {
		if isInOrder(ordering, pages) {
			sum += pages[len(pages)/2]
		}
	}
	return sum
}

func isInOrder(ordering map[int][]int, pages []int) bool {
	for i, n := range pages {
		ord, ok := ordering[n]
		if !ok {
			continue
		}
		for _, x := range ord {
			if slices.Contains(pages[:i], x) {
				return false
			}
		}
	}
	return true
}

func part2(ordering map[int][]int, updates [][]int) int {
	sum := 0
	for _, pages := range updates {
		if !isInOrder(ordering, pages) {
			reorder(ordering, pages)
			sum += pages[len(pages)/2]
		}
	}
	return sum
}

func reorder(ordering map[int][]int, pages []int) {
	for inOrder := false; !inOrder; {
		inOrder = true
		for i, n := range pages {
			ord, ok := ordering[n]
			if !ok {
				continue
			}
			for _, x := range ord {
				if j := slices.Index(pages[:i], x); j >= 0 {
					pages[j], pages[i] = pages[i], pages[j]
					inOrder = false
				}
			}
		}
	}
}
