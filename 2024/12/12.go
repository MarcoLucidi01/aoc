// https://adventofcode.com/2024/day/12
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

func main() {
	grid, err := readGrid(os.Stdin)
	if err != nil {
		panic(err)
	}

	p1, p2 := calcPrices(grid)
	fmt.Println("part1:", p1)
	fmt.Println("part2:", p2)
}

func readGrid(r io.Reader) ([][]byte, error) {
	scanner := bufio.NewScanner(r)

	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(strings.TrimSpace(scanner.Text())))
	}

	return grid, scanner.Err()
}

func calcPrices(grid [][]byte) (int, int) {
	areas := make(map[pos]int)
	perimeters := make(map[pos]int)
	corners := make(map[pos]int)
	visited := make(map[pos]struct{})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			p := pos{i, j}
			if _, ok := visited[p]; ok {
				continue
			}

			a, pe, c := explore(grid, p, grid[i][j], visited)
			areas[p] += a
			perimeters[p] += pe
			corners[p] += c
		}
	}

	p1, p2 := 0, 0
	for p, a := range areas {
		p1 += a * perimeters[p]
		p2 += a * corners[p]
	}
	return p1, p2
}

func explore(grid [][]byte, p pos, plant byte, visited map[pos]struct{}) (int, int, int) {
	if isOutside(grid, p) || grid[p.i][p.j] != plant {
		return 0, 1, 0
	}

	if _, ok := visited[p]; ok {
		return 0, 0, 0
	}
	visited[p] = struct{}{}

	area, peri, corners := 1, 0, countCorners(grid, p, plant)

	next := []pos{
		{p.i - 1, p.j}, // up
		{p.i + 1, p.j}, // down
		{p.i, p.j - 1}, // left
		{p.i, p.j + 1}, // right
	}
	for _, np := range next {
		a, p, c := explore(grid, np, plant, visited)
		area, peri, corners = area+a, peri+p, corners+c
	}
	return area, peri, corners
}

func countCorners(grid [][]byte, p pos, plant byte) int {
	equal := func(p pos) bool { return !isOutside(grid, p) && grid[p.i][p.j] == plant }
	notEqual := func(p pos) bool { return isOutside(grid, p) || grid[p.i][p.j] != plant }

	up := pos{p.i - 1, p.j}
	down := pos{p.i + 1, p.j}
	left := pos{p.i, p.j - 1}
	right := pos{p.i, p.j + 1}

	upLeft := pos{p.i - 1, p.j - 1}
	upRight := pos{p.i - 1, p.j + 1}
	downLeft := pos{p.i + 1, p.j - 1}
	downRight := pos{p.i + 1, p.j + 1}

	c := 0
	if notEqual(up) && notEqual(left) || equal(up) && equal(left) && notEqual(upLeft) {
		c++
	}
	if notEqual(up) && notEqual(right) || equal(up) && equal(right) && notEqual(upRight) {
		c++
	}
	if notEqual(down) && notEqual(left) || equal(down) && equal(left) && notEqual(downLeft) {
		c++
	}
	if notEqual(down) && notEqual(right) || equal(down) && equal(right) && notEqual(downRight) {
		c++
	}
	return c
}

func isOutside(grid [][]byte, p pos) bool {
	return p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[p.i])
}
