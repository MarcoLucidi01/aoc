// https://adventofcode.com/2022/day/1
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cals, err := readSortedCalories(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cals[0])
	fmt.Println(cals[0] + cals[1] + cals[2])
}

func readSortedCalories(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var cals []int
	sum := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			cals = append(cals, sum)
			sum = 0
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		sum += n
	}
	cals = append(cals, sum)
	sort.Sort(sort.Reverse(sort.IntSlice(cals)))
	return cals, scanner.Err()
}
