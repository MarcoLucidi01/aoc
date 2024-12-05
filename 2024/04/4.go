// https://adventofcode.com/2024/day/4
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	grid, err := readGrid(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(grid))
	fmt.Println("part2:", part2(grid))
}

func readGrid(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var grid []string
	for scanner.Scan() {
		grid = append(grid, strings.TrimSpace(scanner.Text()))
	}

	return grid, scanner.Err()
}

func part1(grid []string) int {
	n := 0
	for i, s := range grid {
		for j := range s {
			if grid[i][j] == 'X' {
				n += xmas(grid, i, j)
			}
		}
	}
	return n
}

func xmas(g []string, i, j int) int {
	n := 0
	if strings.HasPrefix(g[i][j:], "XMAS") { // right
		n++
	}
	if j-3 >= 0 && strings.HasPrefix(g[i][j-3:], "SAMX") { // left
		n++
	}
	if i-3 >= 0 && g[i-1][j] == 'M' && g[i-2][j] == 'A' && g[i-3][j] == 'S' { // up
		n++
	}
	if i+3 < len(g) && g[i+1][j] == 'M' && g[i+2][j] == 'A' && g[i+3][j] == 'S' { // down
		n++
	}
	if i+3 < len(g) && j+3 < len(g[i]) && g[i+1][j+1] == 'M' && g[i+2][j+2] == 'A' && g[i+3][j+3] == 'S' { // down right
		n++
	}
	if i+3 < len(g) && j-3 >= 0 && g[i+1][j-1] == 'M' && g[i+2][j-2] == 'A' && g[i+3][j-3] == 'S' { // down left
		n++
	}
	if i-3 >= 0 && j+3 < len(g[i]) && g[i-1][j+1] == 'M' && g[i-2][j+2] == 'A' && g[i-3][j+3] == 'S' { // up right
		n++
	}
	if i-3 >= 0 && j-3 >= 0 && g[i-1][j-1] == 'M' && g[i-2][j-2] == 'A' && g[i-3][j-3] == 'S' { // up left
		n++
	}
	return n
}

func part2(grid []string) int {
	n := 0
	for i, s := range grid {
		for j := range s {
			if grid[i][j] == 'A' && isXMAS(grid, i, j) {
				n++
			}
		}
	}
	return n
}

func isXMAS(g []string, i, j int) bool {
	var a, b [2]byte
	if i-1 >= 0 && j-1 >= 0 { // up left
		a[0] = g[i-1][j-1]
	}
	if i+1 < len(g) && j+1 < len(g[i]) { // down right
		a[1] = g[i+1][j+1]
	}
	if i-1 >= 0 && j+1 < len(g[i]) { // up right
		b[0] = g[i-1][j+1]
	}
	if i+1 < len(g) && j-1 >= 0 { // down left
		b[1] = g[i+1][j-1]
	}

	ms := [2]byte{'M', 'S'}
	sm := [2]byte{'S', 'M'}
	return (a == ms || a == sm) && (b == ms || b == sm)
}
