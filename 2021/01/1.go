// https://adventofcode.com/2021/day/1
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	depths, err := readDepths(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", countIncreased(depths))
	fmt.Printf("part 2: %d\n", countIncreasedThreeSum(depths))
}

func readDepths(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var depths []int
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		if dep == "" {
			continue
		}
		n, err := strconv.Atoi(dep)
		if err != nil {
			continue
		}
		depths = append(depths, n)
	}
	return depths, scanner.Err()
}

func countIncreased(depths []int) int {
	cnt := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			cnt++
		}
	}
	return cnt
}

func countIncreasedThreeSum(depths []int) int {
	cnt := 0
	for i := 0; i+3 < len(depths); i++ {
		a := depths[i] + depths[i+1] + depths[i+2]
		b := depths[i+1] + depths[i+2] + depths[i+3]
		if b > a {
			cnt++
		}
	}
	return cnt
}
