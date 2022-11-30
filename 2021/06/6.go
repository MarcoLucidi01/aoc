// https://adventofcode.com/2021/day/6
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fishes, err := readFishes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", simulateDays(fishes, 80))
	fmt.Printf("part 2: %d\n", simulateDays(fishes, 256))
}

func readFishes() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var fishes []int
	for _, s := range strings.Split(scanner.Text(), ",") {
		if fish, err := strconv.Atoi(s); err == nil {
			fishes = append(fishes, fish)
		}
	}
	return fishes, scanner.Err()
}

func simulateDays(fishes []int, days int) uint64 {
	sim := make([]uint64, 9)
	for _, fish := range fishes {
		sim[fish]++
	}
	for day := 0; day < days; day++ {
		newFishes := sim[0]
		sim = sim[1:]
		sim = append(sim, newFishes)
		sim[6] += newFishes
	}
	nFishes := uint64(0)
	for _, n := range sim {
		nFishes += n
	}
	return nFishes
}
