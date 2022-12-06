// https://adventofcode.com/2022/day/6
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(firstMarker(string(input), 4))
	fmt.Println(firstMarker(string(input), 14))
}

func firstMarker(s string, n int) int {
	lastn := make([]rune, n)
	for i, r := range s {
		lastn[i%n] = r
		if i+1 >= n && !hasDuplicates(lastn) {
			return i + 1
		}
	}
	return 0
}

func hasDuplicates(s []rune) bool {
	set := map[rune]struct{}{}
	for _, r := range s {
		if _, ok := set[r]; ok {
			return true
		}
		set[r] = struct{}{}
	}
	return false
}
