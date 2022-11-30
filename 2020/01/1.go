// https://adventofcode.com/2020/day/1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums, err := readNumbers()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(nums)

	n, x, _ := part1(nums, 2020)
	fmt.Printf("part 1: %d * %d = %d\n", n, x, n*x)

	n, x, z, _ := part2(nums, 2020)
	fmt.Printf("part 2: %d * %d * %d = %d\n", n, x, z, n*x*z)
}

func readNumbers() ([]int, error) {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, scanner.Err()
}

func part1(nums []int, target int) (int, int, bool) {
	j := -1
	k := -1
	for i, n := range nums {
		x := target - n
		if z := sort.SearchInts(nums, x); z < len(nums) && nums[z] == x {
			j = i
			k = z
			break
		}
	}
	if j == -1 || k == -1 {
		return -1, -1, false
	}
	return nums[j], nums[k], true
}

func part2(nums []int, target int) (int, int, int, bool) {
	for _, n := range nums {
		t := target - n
		if x, z, ok := part1(nums, t); ok {
			return n, x, z, true
		}
	}
	return -1, -1, -1, false
}
