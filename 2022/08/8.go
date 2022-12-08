// https://adventofcode.com/2022/day/8
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	grid, err := readTreeGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}

func readTreeGrid(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	var grid [][]int
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		var row []int
		for _, r := range s {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}
	return grid, scanner.Err()
}

func part1(grid [][]int) int {
	nvisible := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if isVisible, _ := rateTree(grid, y, x); isVisible {
				nvisible++
			}
		}
	}
	return nvisible + (len(grid)-1)*2 + (len(grid[0])-1)*2
}

func part2(grid [][]int) int {
	var scores []int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			_, score := rateTree(grid, y, x)
			scores = append(scores, score)
		}
	}

	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}

func rateTree(grid [][]int, y, x int) (bool, int) {
	h := grid[y][x]
	scenicScore := 1
	isVisible := true

	// left
	viewDist, isHidden := 0, false
	for i := x - 1; i >= 0; i-- {
		viewDist++
		if grid[y][i] >= h {
			isHidden = true
			break
		}
	}
	scenicScore *= viewDist
	isVisible = !isHidden

	// right
	viewDist, isHidden = 0, false
	for i := x + 1; i < len(grid[y]); i++ {
		viewDist++
		if grid[y][i] >= h {
			isHidden = true
			break
		}
	}
	scenicScore *= viewDist
	isVisible = isVisible || !isHidden

	// top
	viewDist, isHidden = 0, false
	for i := y - 1; i >= 0; i-- {
		viewDist++
		if grid[i][x] >= h {
			isHidden = true
			break
		}
	}
	scenicScore *= viewDist
	isVisible = isVisible || !isHidden

	// bottom
	viewDist, isHidden = 0, false
	for i := y + 1; i < len(grid); i++ {
		viewDist++
		if grid[i][x] >= h {
			isHidden = true
			break
		}
	}
	scenicScore *= viewDist
	isVisible = isVisible || !isHidden

	return isVisible, scenicScore
}
