// https://adventofcode.com/2024/day/6
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

type dir pos

var (
	Up    = dir{-1, 0}
	Down  = dir{1, 0}
	Left  = dir{0, -1}
	Right = dir{0, 1}

	nextDir = map[dir]dir{
		Up:    Right,
		Right: Down,
		Down:  Left,
		Left:  Up,
	}
)

func main() {
	grid, start, err := readGrid(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(grid, start))
	fmt.Println("part2:", part2(grid, start))
}

func readGrid(r io.Reader) ([][]byte, pos, error) {
	scanner := bufio.NewScanner(r)

	var grid [][]byte
	var start pos
	for i := 0; scanner.Scan(); i++ {
		row := []byte(strings.TrimSpace(scanner.Text()))
		grid = append(grid, row)
		if j := bytes.IndexRune(row, '^'); j >= 0 {
			start = pos{i, j}
		}
	}

	return grid, start, scanner.Err()
}

func part1(grid [][]byte, start pos) int {
	visited := make(map[pos]struct{})
	visited[start] = struct{}{}

	p, dir := start, Up
	for {
		np := pos{p.i + dir.i, p.j + dir.j}

		if isOutside(grid, np) {
			break
		}

		if grid[np.i][np.j] == '#' {
			dir = nextDir[dir]
			continue
		}

		p = np
		visited[p] = struct{}{}
	}

	return len(visited)
}

func isOutside(grid [][]byte, p pos) bool {
	return p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[p.i])
}

func part2(grid [][]byte, start pos) int {
	obstacles := make(map[pos]struct{})

	p, dir := start, Up
	for {
		np := pos{p.i + dir.i, p.j + dir.j}

		if isOutside(grid, np) {
			break
		}

		if grid[np.i][np.j] == '#' {
			dir = nextDir[dir]
			continue
		}

		p = np

		if np == start {
			continue
		}

		if _, ok := obstacles[np]; ok {
			continue
		}

		grid[np.i][np.j] = '#'
		if isLoop(grid, start) {
			obstacles[np] = struct{}{}
		}
		grid[np.i][np.j] = '.'
	}

	return len(obstacles)
}

func isLoop(grid [][]byte, start pos) bool {
	type posdir struct {
		pos
		dir
	}

	visited := make(map[posdir]int)
	visited[posdir{start, Up}]++

	p, dir := start, Up
	for {
		np := pos{p.i + dir.i, p.j + dir.j}

		if isOutside(grid, np) {
			break
		}

		if grid[np.i][np.j] == '#' {
			dir = nextDir[dir]
			continue
		}

		p = np

		visited[posdir{p, dir}]++
		if visited[posdir{p, dir}] > 1 {
			return true
		}
	}

	return false
}
