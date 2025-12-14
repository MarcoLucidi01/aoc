// https://adventofcode.com/2025/day/3
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	banks, err := readBanks(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d\n", part1(banks))
	fmt.Printf("part2: %d\n", part2(banks))
}

func readBanks(r io.Reader) ([][]int, error) {
	var banks [][]int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var bank []int
		for _, r := range scanner.Text() {
			bank = append(bank, int(r-'0'))
		}
		banks = append(banks, bank)
	}

	return banks, scanner.Err()
}

func part1(banks [][]int) int {
	sum := 0
	for _, bank := range banks {
		max := make(map[int]int)
		for i, n := range bank {
			for j := i + 1; j < len(bank); j++ {
				if bank[j] > max[n] {
					max[n] = bank[j]
				}
			}
		}
		m := 0
		for a, b := range max {
			if n := a*10 + b; n > m {
				m = n
			}
		}
		sum += m
	}

	return sum
}

func part2(banks [][]int) int {
	sum := 0
	for _, bank := range banks {
		var m [12]byte
		y := 0
		for x := 11; x >= 0; x-- {
			maxi := 0
			max := 0
			for i := y; i < len(bank)-x; i++ {
				if bank[i] > max {
					maxi = i
					max = bank[i]
				}
			}
			m[-(x - 11)] = byte('0' + max)
			y = maxi + 1
		}
		n, _ := strconv.Atoi(string(m[:]))
		sum += n
	}

	return sum
}
