// https://adventofcode.com/2021/day/9
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	heatmap, err := readHeatmap()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(heatmap))
	fmt.Printf("part 2: %d\n", partTwo(heatmap))
}

func readHeatmap() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var heatmap [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		heatmap = append(heatmap, row)
	}
	return heatmap, scanner.Err()
}

func partOne(heatmap [][]int) int {
	risk := 0
	for y := 0; y < len(heatmap); y++ {
		for x := 0; x < len(heatmap[y]); x++ {
			if isLowPoint(heatmap, x, y) {
				risk += heatmap[y][x] + 1
			}
		}
	}
	return risk
}

func isLowPoint(heatmap [][]int, x, y int) bool {
	switch h := heatmap[y][x]; {
	case x-1 >= 0 && h >= heatmap[y][x-1]:
		return false
	case x+1 < len(heatmap[y]) && h >= heatmap[y][x+1]:
		return false
	case y-1 >= 0 && h >= heatmap[y-1][x]:
		return false
	case y+1 < len(heatmap) && h >= heatmap[y+1][x]:
		return false
	}
	return true
}

func partTwo(heatmap [][]int) int {
	var bSizes []int
	seen := make(map[string]bool)
	for y := 0; y < len(heatmap); y++ {
		for x := 0; x < len(heatmap[y]); x++ {
			if isLowPoint(heatmap, x, y) {
				bSizes = append(bSizes, findBasinSize(heatmap, seen, x, y))
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(bSizes)))
	return bSizes[0] * bSizes[1] * bSizes[2]
}

func findBasinSize(heatmap [][]int, seen map[string]bool, x, y int) int {
	if seen[str(x, y)] || y < 0 || y >= len(heatmap) || x < 0 || x >= len(heatmap[y]) || heatmap[y][x] == 9 {
		return 0
	}
	seen[str(x, y)] = true
	size := findBasinSize(heatmap, seen, x-1, y)
	size += findBasinSize(heatmap, seen, x+1, y)
	size += findBasinSize(heatmap, seen, x, y-1)
	size += findBasinSize(heatmap, seen, x, y+1)
	return size + 1
}

func str(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
