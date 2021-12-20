// https://adventofcode.com/2021/day/20
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	algo, image, err := readAlgoAndImage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", countLit(enhance(algo, image, 2)))
	fmt.Printf("part 2: %d\n", countLit(enhance(algo, image, 50)))
}

func readAlgoAndImage() ([]byte, [][]byte, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	algo := readRow(strings.TrimSpace(scanner.Text()))
	var image [][]byte
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		image = append(image, readRow(line))
	}
	return algo, image, scanner.Err()
}

func readRow(line string) []byte {
	var row []byte
	for _, r := range line {
		pix := byte(1)
		if r == '.' {
			pix = 0
		}
		row = append(row, pix)
	}
	return row
}

func countLit(image [][]byte) int {
	cnt := 0
	for _, row := range image {
		for _, pix := range row {
			cnt += int(pix)
		}
	}
	return cnt
}

func enhance(algo []byte, image [][]byte, nsteps int) [][]byte {
	for step := 0; step < nsteps; step++ {
		image = enhanceStep(algo, image, step)
	}
	return image
}

func enhanceStep(algo []byte, image [][]byte, step int) [][]byte {
	w, h := len(image[0]), len(image)
	var output [][]byte
	for y := -1; y < h+1; y++ {
		var row []byte
		for x := -1; x < w+1; x++ {
			adj := []int{-1, -1, 0, -1, +1, -1, -1, 0, 0, 0, +1, 0, -1, +1, 0, +1, +1, +1}
			index := uint(0)
			for i := 0; i < len(adj); i += 2 {
				index <<= 1
				x1, y1 := x+adj[i], y+adj[i+1]
				if x1 >= 0 && x1 < w && y1 >= 0 && y1 < h {
					index |= uint(image[y1][x1])
				} else if algo[0] == 1 && step%2 == 1 {
					index |= 1
				}
			}
			row = append(row, algo[index])
		}
		output = append(output, row)
	}
	return output
}
