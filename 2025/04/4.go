// https://adventofcode.com/2025/day/4
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	grid, _ := readGrid(os.Stdin)

	fmt.Printf("part1: %d\n", part1(grid))
	fmt.Printf("part2: %d\n", part2(grid))
}

type pos struct {
	y, x int
}

func readGrid(r io.Reader) (map[pos]byte, error) {
	grid := make(map[pos]byte)
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		for j, r := range scanner.Text() {
			grid[pos{i, j}] = byte(r)
		}
	}
	return grid, scanner.Err()
}

func part1(grid map[pos]byte) int {
	x := []int{-1, +0, +1, +1, +1, +0, -1, -1}
	y := []int{+1, +1, +1, +0, -1, -1, -1, +0}
	tot := 0
	for p, v := range grid {
		if v != '@' {
			continue
		}
		n := 0
		for i := 0; i < len(x) && n < 4; i++ {
			if grid[pos{p.y + y[i], p.x + x[i]}] == '@' {
				n++
			}
		}
		if n < 4 {
			tot++
		}
	}

	return tot
}

func part2(grid map[pos]byte) int {
	n := 0
	for {
		acc := accessible(grid)
		if len(acc) == 0 {
			break
		}
		n += len(acc)
		for _, a := range acc {
			delete(grid, a)
		}
	}
	return n
}

func accessible(grid map[pos]byte) []pos {
	x := []int{-1, +0, +1, +1, +1, +0, -1, -1}
	y := []int{+1, +1, +1, +0, -1, -1, -1, +0}
	var ret []pos
	for p, v := range grid {
		if v != '@' {
			continue
		}
		n := 0
		for i := 0; i < len(x) && n < 4; i++ {
			if grid[pos{p.y + y[i], p.x + x[i]}] == '@' {
				n++
			}
		}
		if n < 4 {
			ret = append(ret, p)
		}
	}

	return ret
}
