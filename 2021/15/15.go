// https://adventofcode.com/2021/day/15
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	riskmap, err := readRiskmap()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", findLowestRiskPath(riskmap))
	fmt.Printf("part 2: %d\n", findLowestRiskPath(enlarge(riskmap, 5)))
}

func readRiskmap() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var riskmap [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		riskmap = append(riskmap, row)
	}
	return riskmap, scanner.Err()
}

type point struct {
	x, y int
}

// Dijkstra
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Practical_optimizations_and_infinite_graphs
func findLowestRiskPath(riskmap [][]int) int {
	w, h := len(riskmap[0]), len(riskmap)
	lowrisk := make(map[point]int)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			lowrisk[point{x, y}] = math.MaxInt32
		}
	}

	start := point{0, 0}
	end := point{w - 1, h - 1}
	lowrisk[start] = 0
	pq := PriorityQueue{&Item{point: start, priority: 0}}
	heap.Init(&pq)

	seen := make(map[point]bool)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		p, r := item.point, item.priority
		if p == end {
			return r
		}
		seen[p] = true
		for _, a := range adjacents(p, w, h) {
			if seen[a] {
				continue
			}
			if newr := r + riskmap[a.y][a.x]; newr < lowrisk[a] {
				lowrisk[a] = newr
				heap.Push(&pq, &Item{point: a, priority: newr})
			}
		}
	}
	return -1
}

func adjacents(p point, w, h int) []point {
	a := []int{-1, 0, +1, 0, 0, -1, 0, +1}
	var adj []point
	for i := 0; i < len(a); i += 2 {
		x1, y1 := p.x+a[i], p.y+a[i+1]
		if x1 >= 0 && x1 < w && y1 >= 0 && y1 < h {
			adj = append(adj, point{x1, y1})
		}
	}
	return adj
}

func enlarge(riskmap [][]int, by int) [][]int {
	w, h := len(riskmap[0]), len(riskmap)
	big := make([][]int, h*by)
	for i, _ := range big {
		big[i] = make([]int, w*by)
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			for k := 0; k < by; k++ {
				for z := 0; z < by; z++ {
					x1, y1 := k*w+x, z*h+y
					big[y1][x1] = riskmap[y][x] + k + z
					if big[y1][x1] > 9 {
						big[y1][x1] = big[y1][x1]%10 + 1
					}
				}
			}
		}
	}
	return big
}

// https://pkg.go.dev/container/heap#example-package-PriorityQueue
type Item struct {
	point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
