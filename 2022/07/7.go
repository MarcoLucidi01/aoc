// https://adventofcode.com/2022/day/7
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type fsnode struct {
	parent   *fsnode
	children map[string]*fsnode
	size     int
}

func main() {
	root, err := buildFsTree(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sumDirsAtMost(root, 100000))
	fmt.Println(part2(root))
}

func buildFsTree(r io.Reader) (*fsnode, error) {
	scanner := bufio.NewScanner(r)

	root := &fsnode{children: map[string]*fsnode{}}
	for cwd := root; scanner.Scan(); {
		switch line := strings.Fields(strings.TrimSpace(scanner.Text())); line[0] {
		case "$":
			if line[1] != "cd" {
				continue
			}

			switch dest := line[2]; dest {
			case "/":
				cwd = root
			case "..":
				if cwd.parent != nil {
					cwd = cwd.parent
				}
			default:
				for name, d := range cwd.children {
					if name == dest {
						cwd = d
						break
					}
				}
			}

		case "dir":
			if _, ok := cwd.children[line[1]]; !ok {
				cwd.children[line[1]] = &fsnode{parent: cwd, children: map[string]*fsnode{}}
			}

		default:
			if fsize, err := strconv.Atoi(line[0]); err == nil {
				cwd.children[line[1]] = &fsnode{parent: cwd, size: fsize}
			}
		}
	}

	updateDirSizes(root)
	return root, scanner.Err()
}

func updateDirSizes(f *fsnode) int {
	for _, c := range f.children {
		f.size += updateDirSizes(c)
	}
	return f.size
}

func sumDirsAtMost(f *fsnode, n int) int {
	sum := 0
	if len(f.children) == 0 {
		return sum // file or empty directory
	}

	if f.size < n {
		sum += f.size
	}
	for _, c := range f.children {
		sum += sumDirsAtMost(c, n)
	}

	return sum
}

func part2(root *fsnode) int {
	freeSpace := 70000000 - root.size
	spaceNeeded := 30000000 - freeSpace

	dirs := findDirsBigEnough(root, spaceNeeded)
	if len(dirs) == 0 {
		return 0
	}

	min := dirs[0]
	for _, d := range dirs[1:] {
		if d.size < min.size {
			min = d
		}
	}
	return min.size
}

func findDirsBigEnough(f *fsnode, n int) []*fsnode {
	var dirs []*fsnode
	if len(f.children) == 0 {
		return dirs // file or empty directory
	}

	if f.size >= n {
		dirs = append(dirs, f)
	}
	for _, c := range f.children {
		dirs = append(dirs, findDirsBigEnough(c, n)...)
	}

	return dirs
}
