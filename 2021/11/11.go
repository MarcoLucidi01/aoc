// https://adventofcode.com/2021/day/11
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxEnergy = 9

func main() {
	grid, err := readGrid()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(copyGrid(grid), 100))
	fmt.Printf("part 2: %d\n", partTwo(copyGrid(grid)))
}

func readGrid() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}
	return grid, scanner.Err()
}

func copyGrid(grid [][]int) [][]int {
	cp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		cp[i] = append(cp[i], grid[i]...)
	}
	return cp
}

func partOne(grid [][]int, nsteps int) int {
	nflash := 0
	for i := 0; i < nsteps; i++ {
		nflash += step(grid)
	}
	return nflash
}

func step(grid [][]int) int {
	var adj []int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x]++
			if grid[y][x] > maxEnergy {
				adj = appendAdjacent(adj, x, y, len(grid[0]), len(grid))
			}
		}
	}
	for len(adj) > 0 {
		x, y := adj[len(adj)-2], adj[len(adj)-1]
		adj = adj[:len(adj)-2]
		if grid[y][x] > maxEnergy {
			continue // already flashed
		}
		grid[y][x]++
		if grid[y][x] > maxEnergy {
			adj = appendAdjacent(adj, x, y, len(grid[0]), len(grid))
		}
	}
	nflash := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] > maxEnergy {
				nflash++
				grid[y][x] = 0
			}
		}
	}
	return nflash
}

func appendAdjacent(adj []int, x, y, w, h int) []int {
	a := []int{-1, 0, +1, 0, 0, -1, 0, +1, -1, -1, -1, +1, +1, -1, +1, +1}
	for i := 0; i < len(a); i += 2 {
		x1, y1 := x+a[i], y+a[i+1]
		if x1 >= 0 && x1 < w && y1 >= 0 && y1 < h {
			adj = append(adj, x1, y1)
		}
	}
	return adj
}

func partTwo(grid [][]int) int {
	var i int
	for i = 1; step(grid) != len(grid[0])*len(grid); i++ {
	}
	return i
}
