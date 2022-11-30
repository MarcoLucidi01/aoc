// https://adventofcode.com/2020/day/4
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	nvalidPart1, nvalidPart2, err := countValidPassports(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", nvalidPart1)
	fmt.Printf("part 2: %d\n", nvalidPart2)
}

func countValidPassports(r io.Reader) (int, int, error) {
	nvalidPart1 := 0
	nvalidPart2 := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		passport := readPassportFields(scanner)
		nrequiredPart1 := 0
		nrequiredPart2 := 0
		for k, v := range passport {
			isValidFunc, ok := required[k]
			if !ok {
				continue // not required, skip
			}
			nrequiredPart1++
			if isValidFunc(v) {
				nrequiredPart2++
			}
		}
		if nrequiredPart1 == len(required) {
			nvalidPart1++
		}
		if nrequiredPart2 == len(required) {
			nvalidPart2++
		}
	}
	return nvalidPart1, nvalidPart2, scanner.Err()
}

func readPassportFields(scanner *bufio.Scanner) map[string]string {
	fields := make(map[string]string)
	for {
		line := scanner.Text()
		if line == "" {
			break
		}
		pairs := strings.Split(line, " ")
		for _, p := range pairs {
			pair := strings.Split(p, ":")
			if len(pair) != 2 {
				continue
			}
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])
			if len(key) == 0 || len(value) == 0 {
				continue
			}
			fields[key] = value
		}
		if !scanner.Scan() {
			break
		}
	}
	return fields
}

var (
	reYr  = regexp.MustCompile(`^\d{4}$`)
	reHgt = regexp.MustCompile(`^(\d+)(cm|in)$`)
	reHcl = regexp.MustCompile(`^#([0-9]|[a-f]){6}$`)
	reEcl = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	rePid = regexp.MustCompile(`^\d{9}$`)

	required = map[string]func(string) bool{
		"byr": func(v string) bool { return isValidYear(v, 1920, 2002) },
		"iyr": func(v string) bool { return isValidYear(v, 2010, 2020) },
		"eyr": func(v string) bool { return isValidYear(v, 2020, 2030) },
		"hgt": func(v string) bool {
			match := reHgt.FindStringSubmatch(v)
			if len(match) != 3 || match[0] != v {
				return false
			}
			n, err := strconv.Atoi(match[1])
			if err != nil {
				return false
			}
			switch match[2] {
			case "cm":
				if n < 150 || n > 193 {
					return false
				}
			case "in":
				if n < 59 || n > 76 {
					return false
				}
			default:
				return false
			}
			return true
		},
		"hcl": func(v string) bool { return reHcl.MatchString(v) },
		"ecl": func(v string) bool { return reEcl.MatchString(v) },
		"pid": func(v string) bool { return rePid.MatchString(v) },
	}
)

func isValidYear(v string, min, max int) bool {
	if !reYr.MatchString(v) {
		return false
	}
	year, err := strconv.Atoi(v)
	return err == nil && year >= min && year <= max
}
