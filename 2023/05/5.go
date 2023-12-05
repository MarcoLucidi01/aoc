// https://adventofcode.com/2023/day/5
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	seeds, maps, err := parseAlmanac(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", part1(seeds, maps))
	fmt.Println("part2:", part2(seeds, maps))
}

func parseAlmanac(r io.Reader) ([]int, [][][3]int, error) {
	scanner := bufio.NewScanner(r)

	var seeds []int
	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())
	i := strings.IndexRune(s, ':')
	for _, f := range strings.Fields(s[i+1:]) {
		n, err := strconv.Atoi(strings.TrimSpace(f))
		if err != nil {
			return nil, nil, err
		}
		seeds = append(seeds, n)
	}

	var maps [][][3]int
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" || s[len(s)-1] != ':' {
			continue
		}

		var m [][3]int
		for scanner.Scan() {
			s := strings.TrimSpace(scanner.Text())
			if s == "" {
				break
			}

			var line [3]int
			for i, f := range strings.Fields(s) {
				n, err := strconv.Atoi(strings.TrimSpace(f))
				if err != nil {
					return nil, nil, err
				}
				line[i] = n
			}
			m = append(m, line)
		}
		maps = append(maps, m)
	}

	return seeds, maps, scanner.Err()
}

func part1(seeds []int, maps [][][3]int) int {
	low := math.MaxInt32
	for _, seed := range seeds {
		for k, m := range maps {
			for _, r := range m {
				if seed >= r[1] && seed < r[1]+r[2] {
					seed = r[0] + (seed - r[1])
					break
				}
			}
			if k == len(maps)-1 && seed < low {
				low = seed
			}
		}
	}
	return low
}

// this is so slow, but it works :)
func part2(seeds []int, maps [][][3]int) int {
	low := math.MaxInt32
	for i := 0; i < len(seeds); i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			seed := seeds[i] + j
			for k, m := range maps {
				for _, r := range m {
					if seed >= r[1] && seed < r[1]+r[2] {
						seed = r[0] + (seed - r[1])
						break
					}
				}
				if k == len(maps)-1 && seed < low {
					low = seed
				}
			}
		}
	}
	return low
}
