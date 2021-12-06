/* https://adventofcode.com/2020/day/9 */

#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void    readnumbers(int **, size_t *);
int     readint(void);
int     partone(int *, size_t, size_t);
bool    isvalid(int *, size_t, int);
int     parttwo(int *, size_t, int);
void   *xrealloc(void *, size_t);
void   *xmalloc(size_t);
void    die(const char *);

enum {
        PREAMBLELEN = 25,  /* default for the input */
        BUFSIZE     = 256, /* for reading numbers */
};

int main(int argc, char **argv)
{
        int prelen = PREAMBLELEN;
        if (argc > 1) {
                prelen = atoi(argv[1]);
                if (prelen < 2)
                        die("invalid preamble length, should be at least 2");
        }

        int *numbers = NULL;
        size_t len = 0;
        readnumbers(&numbers, &len);
        if ((int)len <= prelen) {
                free(numbers);
                die("too few numbers");
        }

        int firstnotvalid = partone(numbers, len, prelen);
        printf("part 1: %d\n", firstnotvalid);
        printf("part 2: %d\n", parttwo(numbers, len, firstnotvalid));

        free(numbers);
}

void readnumbers(int **numbers, size_t *len)
{
        *numbers = NULL;
        *len = 0;
        size_t cap = 0;

        int n;
        while ((n = readint()) != INT_MIN) {
                if (*len == cap) {
                        cap += cap / 2 + 10;
                        *numbers = xrealloc(*numbers, sizeof(**numbers) * cap);
                }
                (*numbers)[(*len)++] = n;
        }
}

int readint(void)
{
        char buf[BUFSIZE];
        if (fgets(buf, sizeof(buf), stdin) == NULL) {
                if (ferror(stdin))
                        die("fgets: read error");
                if (feof(stdin))
                        return INT_MIN;
        }
        return atoi(buf);
}

int partone(int *numbers, size_t len, size_t prelen)
{
        int *preamble = xmalloc(sizeof(*preamble) * prelen);
        memcpy(preamble, numbers, sizeof(*preamble) * prelen);

        int n = INT_MIN;
        for (size_t i = prelen; i < len; i++) {
                n = numbers[i];
                if ( ! isvalid(preamble, prelen, n))
                        break;
                memmove(preamble, preamble + 1, sizeof(*preamble) * (prelen - 1));
                preamble[prelen - 1] = n;
        }
        free(preamble);
        return n;
}

bool isvalid(int *preamble, size_t prelen, int n)
{
        for (size_t i = 0; i < prelen; i++) {
                for (size_t j = 0; j < prelen; j++) {
                        if (preamble[i] == preamble[j])
                                continue;
                        if (preamble[i] + preamble[j] == n)
                                return true;
                }
        }
        return false;
}

int parttwo(int *numbers, size_t len, int invalid)
{
        size_t a = 0;
        size_t b = 0;
        for (size_t i = 0; i < len; i++) {
                a = i;
                long sum = numbers[i];
                for (size_t j = i+1; j < len; j++) {
                        sum += numbers[j];
                        if (sum > invalid)
                                break;
                        if (sum == invalid) {
                                b = j;
                                goto out;
                        }
                }
        }
out:;
        int min = numbers[a];
        for (size_t i = a; i < b; i++)
                if (numbers[i] < min)
                        min = numbers[i];

        int max = min;
        for (size_t i = a; i < b; i++)
                if (numbers[i] > max)
                        max = numbers[i];

        return min + max;
}

void *xrealloc(void *p, size_t n)
{
        p = realloc(p, n);
        if (p == NULL && n != 0)
                die("realloc: out of memory");
        return p;
}

void *xmalloc(size_t n)
{
        return xrealloc(NULL, n);
}

void die(const char *s)
{
        fprintf(stderr, "error: %s\n", s);
        exit(EXIT_FAILURE);
}
