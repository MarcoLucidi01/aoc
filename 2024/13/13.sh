#!/bin/sh

# https://adventofcode.com/2024/day/13

# https://en.wikipedia.org/wiki/Cramer%27s_rule
grep -Eo '[0-9]+' \
 | paste - - - - - - \
 | awk '
   {
       ax=$1; ay=$2; bx=$3; by=$4; px=$5; py=$6;

       A = (px*by - py*bx) / (ax*by - ay*bx)
       B = (ax*py - ay*px) / (ax*by - ay*bx)
       if (A%1 == 0 && B%1 == 0) {
           p1 += 3*A + B
       }

       px += 10000000000000
       py += 10000000000000
       A = (px*by - py*bx) / (ax*by - ay*bx)
       B = (ax*py - ay*px) / (ax*by - ay*bx)
       if (A%1 == 0 && B%1 == 0) {
           p2 += 3*A + B
       }
   }
   END { printf("part1: %d\npart2: %d\n", p1, p2) }
'
