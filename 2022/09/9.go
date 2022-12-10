// https://adventofcode.com/2022/day/9
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

type pos struct {
	x, y int
}

type move struct {
	dir   string
	steps int
}

func main() {
	moves, err := readMoves(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(simulate(moves, 2))
	fmt.Println(simulate(moves, 10))
}

func readMoves(r io.Reader) ([]move, error) {
	scanner := bufio.NewScanner(r)
	var moves []move
	for scanner.Scan() {
		fields := strings.Fields(strings.TrimSpace(scanner.Text()))
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid move %v", fields)
		}
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		moves = append(moves, move{dir: fields[0], steps: n})
	}
	return moves, scanner.Err()
}

func simulate(moves []move, n int) int {
	var rope []pos
	for i := 0; i < n; i++ {
		rope = append(rope, pos{x: 0, y: 0})
	}

	visited := map[pos]struct{}{}
	visited[rope[n-1]] = struct{}{}

	increments := map[string]pos{
		"U": pos{x: 0, y: 1},
		"D": pos{x: 0, y: -1},
		"L": pos{x: -1, y: 0},
		"R": pos{x: 1, y: 0},
	}

	for _, m := range moves {
		incr := increments[m.dir]
		for i := 0; i < m.steps; i++ {
			rope[0].x += incr.x
			rope[0].y += incr.y
			for j := 1; j < len(rope); j++ {
				rope[j] = nextPos(rope[j-1], rope[j])
			}
			visited[rope[n-1]] = struct{}{}
		}
	}

	return len(visited)
}

// who said bruteforce? ehehe
func nextPos(h, t pos) pos {
	hNear := []pos{
		h,
		{x: h.x, y: h.y + 1},
		{x: h.x, y: h.y - 1},
		{x: h.x - 1, y: h.y},
		{x: h.x + 1, y: h.y},

		{x: h.x - 1, y: h.y - 1},
		{x: h.x - 1, y: h.y + 1},
		{x: h.x + 1, y: h.y - 1},
		{x: h.x + 1, y: h.y + 1},
	}
	for _, hn := range hNear {
		if t == hn {
			return t // tail is already close to the head
		}
	}

	tNear := []pos{ // try every new position the tail can be
		{x: t.x, y: t.y + 1},
		{x: t.x, y: t.y - 1},
		{x: t.x - 1, y: t.y},
		{x: t.x + 1, y: t.y},

		{x: t.x - 1, y: t.y - 1},
		{x: t.x - 1, y: t.y + 1},
		{x: t.x + 1, y: t.y - 1},
		{x: t.x + 1, y: t.y + 1},
	}
	for i, tn := range tNear {
		// consider up/down/left/right new positions only if tail and
		// head are on the same row/column
		if (i == 0 || i == 1) && h.x != t.x {
			continue
		}
		if (i == 2 || i == 3) && h.y != t.y {
			continue
		}
		for _, hn := range hNear {
			if tn == hn {
				return tn
			}
		}
	}
	return t
}
