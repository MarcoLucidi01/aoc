#!/bin/awk -f
# https://adventofcode.com/2021/day/2

$1=="forward" { h += $2; d2 += $2*d }
$1=="down"    { d += $2 }
$1=="up"      { d -= $2 }
END           { printf "part 1: %d\npart 2: %d\n", h*d, h*d2 }
