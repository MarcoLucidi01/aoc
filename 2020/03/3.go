package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	// remove last line if it's empty (i.e. when input ends with '\n')
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	ntrees := countTrees(lines, 3, 1)
	fmt.Printf("part 1: %d\n", ntrees)

	ntrees11 := countTrees(lines, 1, 1)
	ntrees31 := countTrees(lines, 3, 1)
	ntrees51 := countTrees(lines, 5, 1)
	ntrees71 := countTrees(lines, 7, 1)
	ntrees12 := countTrees(lines, 1, 2)
	fmt.Printf("part 2: %d\n", ntrees11*ntrees31*ntrees51*ntrees71*ntrees12)
}

func countTrees(lines []string, right, down int) int {
	ntrees := 0
	for x, y := 0, 0; y < len(lines); y += down {
		if lines[y][x] == '#' {
			ntrees++
		}
		x = (x + right) % len(lines[y])
	}
	return ntrees
}
