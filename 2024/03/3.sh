#!/bin/sh

# https://adventofcode.com/2024/day/3

grep -Eo "mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)" \
 | sed 's/^mul(\([0-9]\+\),\([0-9]\+\))$/\1 \2/' \
 | awk '
   BEGIN      { d=1 }
   /^do\(.*$/ { d=1; next }
   /^don.*$/  { d=0; next }
              { p1 += $1*$2 }
   d          { p2 += $1*$2 }
   END        { printf("part1: %d\npart2: %d\n", p1, p2) }
'
