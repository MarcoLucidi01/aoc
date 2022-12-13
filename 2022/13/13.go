// https://adventofcode.com/2022/day/13
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	packets, err := readPackets(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(packets))
	fmt.Println(part2(packets))
}

func readPackets(r io.Reader) ([][]any, error) {
	scanner := bufio.NewScanner(r)
	var packets [][]any
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}

		var pack []any
		if err := json.Unmarshal([]byte(s), &pack); err != nil {
			return nil, err
		}

		packets = append(packets, pack)
	}

	return packets, scanner.Err()
}

func part1(packets [][]any) int {
	sum := 0
	for i := 0; i < len(packets); i += 2 {
		if compare(packets[i], packets[i+1]) < 0 {
			sum += i/2 + 1
		}
	}
	return sum
}

func part2(packets [][]any) int {
	d1 := []any{[]any{2.0}}
	d2 := []any{[]any{6.0}}
	packets = append(packets, d1, d2)

	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })

	i1 := sort.Search(len(packets), func(i int) bool { return compare(packets[i], d1) >= 0 })
	i2 := sort.Search(len(packets), func(i int) bool { return compare(packets[i], d2) >= 0 })
	return (i1 + 1) * (i2 + 1)
}

func compare(a, b any) int {
	na, okNa := a.(float64) // json numbers are floats
	nb, okNb := b.(float64)
	la, okLa := a.([]any)
	lb, okLb := b.([]any)

	switch {
	case okNa && okNb: // both numbers
		return int(na - nb)

	case okNa && okLb: // number and list
		la = []any{na}
		return compare(la, lb)

	case okNb && okLa: // number and list
		lb = []any{nb}
		return compare(la, lb)

	case okLa && okLb: // both lists
		minLen := len(la)
		if len(lb) < minLen {
			minLen = len(lb)
		}
		for i := 0; i < minLen; i++ {
			if cmp := compare(la[i], lb[i]); cmp != 0 {
				return cmp
			}
		}
		return len(la) - len(lb)
	}

	return 0
}
