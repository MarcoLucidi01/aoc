/*
 * https://adventofcode.com/2021/day/24
 * simplified and refactored C version of MONAD program specific to my input.
 *
 * I solved the puzzle by hand with trial and error observing how the values of
 * x and z change after each step and which digits in the testing number causes
 * the bigger changes.
 *
 * part 1: 69914999975369
 * part 2: 14911675311114
 */

#include <ctype.h>
#include <stdarg.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

enum {
        NDIGITS = 14,
};

static bool ismonad(const int *);
static void die(const char *, ...);

int main(int argc, char **argv)
{
        if (argc != 2)
                die("wrong number of arguments\nusage: %s %d-digits-number", argv[0], NDIGITS);

        const char *number = argv[1];
        int digits[NDIGITS];
        int i;
        for (i = 0; i < NDIGITS; i++) {
                if (number[i] == '\0')
                        die("too few digits, number must have %d digits", NDIGITS);
                if ( ! isdigit(number[i]))
                        die("invalid digit '%c'", number[i]);
                digits[i] = number[i] - '0';
        }
        if (number[i] != '\0')
                die("too much digits, number must have %d digits", NDIGITS);

        exit(ismonad(digits) ? EXIT_SUCCESS : EXIT_FAILURE);
}

static bool ismonad(const int *digits)
{
        int a[NDIGITS] = { 10, 13, 13, -11, 11, -4, 12, 12, 15, -2, -5, -11, -13, -10 };
        int b[NDIGITS] = { 1,   1,  1,  26,  1, 26,  1,  1,  1, 26, 26,  26,  26,  26 };
        int c[NDIGITS] = { 13, 10,  3,   1,  9,  3,  5,  1,  0, 13,  7,  15,  12,   8 };

        long z = 0;
        for (int i = 0; i < NDIGITS; i++) {
                int x = z % 26 + a[i];
                z /= b[i];
                if (x != digits[i])
                        z = z * 26 + digits[i] + c[i];

                printf("digits[%2d]: %d, a: %3d, b: %3d, c: %2d, x: %3d, z: %13ld\n", i, digits[i], a[i], b[i], c[i], x, z);
        }
        return z == 0;
}

static void die(const char *reason, ...)
{
        fprintf(stderr, "error: ");
        va_list ap;
        va_start(ap, reason);
        vfprintf(stderr, reason, ap);
        va_end(ap);
        fprintf(stderr, "\n");

        exit(EXIT_FAILURE);
}
