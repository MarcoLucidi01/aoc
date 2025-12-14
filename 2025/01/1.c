/* https://adventofcode.com/2025/day/1 */

#include <stdio.h>
#include <stdlib.h>

int main(void)
{
    int dial = 50;
    int part1 = 0;
    int part2 = 0;
    char s[256];
    while (fgets(s, sizeof(s), stdin)) {
        int n = atoi(s+1);
        for (int i = 0; i < n; i++) {
            dial += s[0] == 'L' ? -1 : +1;
            if (dial == -1)
                dial = 99;
            else if (dial == 100)
                dial = 0;

            if (dial == 0)
                part2++;
        }
        if (dial == 0)
            part1++;
    }

    printf("part1 %d\npart2 %d\n", part1, part2);
}
