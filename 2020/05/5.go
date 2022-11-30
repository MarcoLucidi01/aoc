// https://adventofcode.com/2020/day/5
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxId := -1
	var seatIds []int
	for scanner.Scan() {
		n := toBinary(scanner.Text())
		id := computeSeatId(n)
		seatIds = append(seatIds, id)
		if id > maxId {
			maxId = id
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Printf("part 1: %d\n", maxId)

	sort.Ints(seatIds)
	myID := -1
	for i, id := range seatIds {
		if i > 0 && id == seatIds[i-1]+2 {
			myID = seatIds[i-1] + 1
			break
		}
	}
	fmt.Printf("part 2: %d\n", myID)
}

func toBinary(seatno string) int {
	bin := strings.ToUpper(seatno)
	bin = strings.ReplaceAll(bin, "B", "1")
	bin = strings.ReplaceAll(bin, "F", "0")
	bin = strings.ReplaceAll(bin, "R", "1")
	bin = strings.ReplaceAll(bin, "L", "0")
	n, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return -1
	}
	return int(n)
}

func computeSeatId(n int) int {
	row := (n & 0b1111111000) >> 3
	col := n & 0b111
	return row*8 + col
}
