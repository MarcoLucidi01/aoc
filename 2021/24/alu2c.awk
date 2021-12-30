#!/bin/awk -f
# https://adventofcode.com/2021/day/24
# translates ALU input instructions into a ready to run-understand-refactor-debug C program.

BEGIN {
        print "#include <stdio.h>"
        print ""
        print "int main(void)"
        print "{"
        print "\t/* CHANGE INPUT HERE */"
        print "\tlong input[] = { 1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9 };"
        print ""
        print "\tlong w = 0;"
        print "\tlong x = 0;"
        print "\tlong y = 0;"
        print "\tlong z = 0;"
        print ""
}

{ printf "\t" }
$1 == "inp" { printf "%s = input[%d];", $2, i++ }
$1 == "add" { printf "%s += %s;", $2, $3 }
$1 == "mul" { printf "%s *= %s;", $2, $3 }
$1 == "div" { printf "%s /= %s;", $2, $3 }
$1 == "mod" { printf "%s %%= %s;", $2, $3 }
$1 == "eql" { printf "%s = %s == %s ? 1 : 0;", $2, $2, $3 }
{ print "" }

END {
        print ""
        print "\tprintf(\"w: %ld, x: %ld, y: %ld, z: %ld\\n\", w, x, y, z);"
        print "}"
}
