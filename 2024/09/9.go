// https://adventofcode.com/2024/day/9
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	diskMap, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	diskMap = bytes.TrimSpace(diskMap)

	fmt.Println("part1:", checksum(defrag1(expand(diskMap))))
	fmt.Println("part2:", checksum(defrag2(expand(diskMap))))
}

func expand(diskMap []byte) []int {
	var buf []int
	id := 0
	for i, b := range diskMap {
		n := int(b - '0')

		if i%2 == 0 {
			for j := 0; j < n; j++ {
				buf = append(buf, id)
			}
			id++
			continue
		}

		for j := 0; j < n; j++ {
			buf = append(buf, -1)
		}
	}
	return buf
}

func defrag1(disk []int) []int {
	j := 0
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}

		for ; disk[j] != -1; j++ {
		}
		if j >= i {
			break
		}

		disk[j], disk[i] = disk[i], disk[j]
	}
	return disk
}

func checksum(disk []int) int {
	sum := 0
	for i, n := range disk {
		if n != -1 {
			sum += i * n
		}
	}
	return sum
}

func defrag2(disk []int) []int {
loop:
	for e1 := len(disk) - 1; e1 >= 0; {
		if disk[e1] == -1 {
			e1--
			continue
		}

		s1 := e1
		for ; s1 >= 0 && disk[s1] == disk[e1]; s1-- {
		}
		s1++
		size := e1 - s1 + 1

		s2 := 0
		e2 := s2
		for {
			for ; e2 < s1 && disk[e2] != -1; e2++ {
			}
			if e2 >= s1 {
				e1 = s1 - 1
				continue loop
			}

			s2 = e2
			for ; e2 < len(disk) && disk[e2] == -1 && e2-s2 < size; e2++ {
			}
			if e2 < len(disk) && e2-s2 >= size {
				break
			}
		}

		for ; s1 <= e1; s1, s2 = s1+1, s2+1 {
			disk[s2], disk[s1] = disk[s1], disk[s2]
		}
		e1 -= size
	}
	return disk
}
