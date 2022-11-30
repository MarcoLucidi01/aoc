// https://adventofcode.com/2020/day/6
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type countFunc func([]string) int

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	fmt.Printf("part 1: %d\n", countYes(groups, anyone))
	fmt.Printf("part 2: %d\n", countYes(groups, everyone))
}

func countYes(groups []string, cfn countFunc) int {
	var counts []int
	for _, group := range groups {
		counts = append(counts, cfn(strings.Split(strings.TrimSpace(group), "\n")))
	}
	sum := 0
	for _, cnt := range counts {
		sum += cnt
	}
	return sum
}

func anyone(group []string) int {
	yes := make(map[rune]bool)
	for _, ans := range group {
		for _, a := range ans {
			yes[a] = true
		}
	}
	return len(yes)
}

func everyone(group []string) int {
	yes := make(map[rune]int)
	for _, ans := range group {
		for _, a := range ans {
			yes[a]++
		}
	}
	cnt := 0
	for _, n := range yes {
		if n == len(group) {
			cnt++
		}
	}
	return cnt
}
