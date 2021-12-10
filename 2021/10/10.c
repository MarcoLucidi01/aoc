// https://adventofcode.com/2021/day/10

#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

static int      cmplong(const void *, const void *);
static void    *xrealloc(void *, size_t);

enum {
        MAXLINESIZE = 4096,
};

static unsigned char closetab[UCHAR_MAX + 1] = {
        ['('] = ')',
        ['['] = ']',
        ['{'] = '}',
        ['<'] = '>',
};

static short pointsp1[UCHAR_MAX + 1] = {
        [')'] = 3,
        [']'] = 57,
        ['}'] = 1197,
        ['>'] = 25137,
};

static short pointsp2[UCHAR_MAX + 1] = {
        [')'] = 1,
        [']'] = 2,
        ['}'] = 3,
        ['>'] = 4,
};

int main(void)
{
        long scorep1 = 0;
        long *scoresp2 = NULL;
        size_t nscoresp2 = 0;
        char *line = xrealloc(NULL, MAXLINESIZE);
        while (fgets(line, MAXLINESIZE, stdin) != NULL) {
                int sp = 0;
                unsigned char *c = (unsigned char *)line;
                for (; *c != '\n' && *c != '\0'; c++) {
                        if (closetab[*c] != 0)
                                line[sp++] = closetab[*c];
                        else if (sp == 0 || *c != line[--sp])
                                break;
                }
                if (*c != '\n' && *c != '\0') { // corrupted line
                        scorep1 += pointsp1[*c];
                        continue;
                }
                long sc = 0;
                while (sp > 0)
                        sc = sc * 5 + pointsp2[(unsigned char)line[--sp]];
                scoresp2 = xrealloc(scoresp2, sizeof(*scoresp2) * (nscoresp2 + 1)); // really bad growth factor eheh
                scoresp2[nscoresp2++] = sc;
        }
        qsort(scoresp2, nscoresp2, sizeof(*scoresp2), cmplong);
        long scorep2 = nscoresp2 > 0 ? scoresp2[nscoresp2 / 2] : 0;

        printf("part 1: %ld\n", scorep1);
        printf("part 2: %ld\n", scorep2);

        free(line);
        free(scoresp2);
}

static int cmplong(const void *pa, const void *pb)
{
        long a = *((long *)pa);
        long b = *((long *)pb);
        return a == b ? 0 : (a > b ? 1 : -1);
}

static void *xrealloc(void *p, size_t n)
{
        p = realloc(p, n);
        if (p == NULL && n != 0) {
                fprintf(stderr, "xrealloc: out of memory");
                exit(EXIT_FAILURE);
        }
        return p;
}
