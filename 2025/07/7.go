// https://adventofcode.com/2025/day/7
package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"strings"
)

func main() {
	start, diagram, err := readDiagram(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d\n", part1(start, maps.Clone(diagram)))
	fmt.Printf("part2: %d\n", part2(start, maps.Clone(diagram)))
}

type pos struct {
	x, y int
}

func readDiagram(r io.Reader) (pos, map[pos]byte, error) {
	scanner := bufio.NewScanner(r)

	var start pos
	diagram := make(map[pos]byte)
	for i := 0; scanner.Scan(); i++ {
		for j, r := range strings.TrimSpace(scanner.Text()) {
			p := pos{i, j}
			diagram[p] = byte(r)
			if r == 'S' {
				start = p
			}
		}
	}

	return start, diagram, scanner.Err()
}

func part1(start pos, diagram map[pos]byte) int {
	q := []pos{start}
	splits := 0
	for len(q) > 0 {
		p := q[len(q)-1]
		q = q[:len(q)-1]

		p = pos{p.x + 1, p.y} // down
		switch diagram[p] {
		case '.':
			diagram[p] = '|'
			q = append(q, p)
		case '^':
			q = append(q,
				pos{p.x, p.y - 1}, // left
				pos{p.x, p.y + 1}, // right
			)
			splits++
		}
	}
	return splits
}

func part2(start pos, diagram map[pos]byte) int {
	return timelines(start, diagram, make(map[pos]int))
}

func timelines(p pos, diagram map[pos]byte, cache map[pos]int) int {
	if n, ok := cache[p]; ok {
		return n
	}

	switch diagram[p] {
	case 'S', '.':
		n := timelines(pos{p.x + 1, p.y}, diagram, cache) // down
		cache[p] = n
		return n
	case '^':
		left := timelines(pos{p.x, p.y - 1}, diagram, cache)
		right := timelines(pos{p.x, p.y + 1}, diagram, cache)
		cache[p] = left + right
		return left + right
	default:
		cache[p] = 1
		return 1
	}
}
