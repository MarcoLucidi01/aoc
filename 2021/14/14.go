// https://adventofcode.com/2021/day/14
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	template, rules, err := readTemplateAndRules()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", solve(template, rules, 10))
	fmt.Printf("part 2: %d\n", solve(template, rules, 40))
}

func readTemplateAndRules() (string, map[string]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	template := scanner.Text()
	rules := make(map[string]string)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		split := strings.Split(line, " -> ")
		rules[split[0]] = split[1]
	}
	return template, rules, scanner.Err()
}

func solve(template string, rules map[string]string, steps int) uint {
	pairs := make(map[string]uint)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		pairs[pair]++
	}

	var nextPairs map[string]uint
	for i := 0; i < steps-1; i++ {
		nextPairs = make(map[string]uint)
		for pair, n := range pairs {
			element := rules[pair]
			nextPairs[string(pair[0])+element] += n
			nextPairs[element+string(pair[1])] += n
		}
		pairs = nextPairs
	}
	polymer := make(map[string]uint)
	for pair, n := range nextPairs {
		element := rules[pair]
		polymer[string(pair[0])+element] += n
	}

	freq := make(map[rune]uint)
	for pair, n := range polymer {
		for _, element := range pair {
			freq[element] += n
		}
	}
	freq[rune(template[len(template)-1])]++

	min, max := ^uint(0), uint(0)
	for _, f := range freq {
		if f > max {
			max = f
		}
		if f < min {
			min = f
		}
	}
	return max - min
}
