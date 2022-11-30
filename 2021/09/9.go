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
	hmap, err := readHeightmap()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(hmap))
	fmt.Printf("part 2: %d\n", partTwo(hmap))
}

func readHeightmap() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var hmap [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		hmap = append(hmap, row)
	}
	return hmap, scanner.Err()
}

func partOne(hmap [][]int) int {
	risk := 0
	for y := 0; y < len(hmap); y++ {
		for x := 0; x < len(hmap[y]); x++ {
			if isLowPoint(hmap, x, y) {
				risk += hmap[y][x] + 1
			}
		}
	}
	return risk
}

func isLowPoint(hmap [][]int, x, y int) bool {
	switch h := hmap[y][x]; {
	case x-1 >= 0 && h >= hmap[y][x-1]:
		return false
	case x+1 < len(hmap[y]) && h >= hmap[y][x+1]:
		return false
	case y-1 >= 0 && h >= hmap[y-1][x]:
		return false
	case y+1 < len(hmap) && h >= hmap[y+1][x]:
		return false
	}
	return true
}

type point struct {
	x, y int
}

func partTwo(hmap [][]int) int {
	var bsizes []int
	seen := make(map[point]bool)
	for y := 0; y < len(hmap); y++ {
		for x := 0; x < len(hmap[y]); x++ {
			if bsize := findBasinSize(hmap, seen, x, y); bsize > 0 {
				bsizes = append(bsizes, bsize)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(bsizes)))
	return bsizes[0] * bsizes[1] * bsizes[2]
}

func findBasinSize(hmap [][]int, seen map[point]bool, x, y int) int {
	if y < 0 || y >= len(hmap) || x < 0 || x >= len(hmap[y]) || hmap[y][x] == 9 || seen[point{x, y}] {
		return 0
	}
	seen[point{x, y}] = true
	size := findBasinSize(hmap, seen, x-1, y)
	size += findBasinSize(hmap, seen, x+1, y)
	size += findBasinSize(hmap, seen, x, y-1)
	size += findBasinSize(hmap, seen, x, y+1)
	return size + 1
}
