/* https://adventofcode.com/2021/day/17 */

#include <stdio.h>
#include <stdlib.h>

enum {
        N = 500,
};

int main(void)
{
        int tx1, tx2, ty1, ty2; /* target */
        if (scanf("target area: x=%d..%d, y=%d..%d", &tx1, &tx2, &ty1, &ty2) != 4) {
                fprintf(stderr, "invalid input\n");
                exit(EXIT_FAILURE);
        }

        int maxheight = 0;
        int ngood = 0;
        for (int i = -N; i < N; i++) { /* stupid brute force, can easily be improved */
                for (int j = -N; j < N; j++) {
                        int px = 0; /* probe position */
                        int py = 0;
                        int vx = i; /* probe velocity */
                        int vy = j;
                        int maxstepheight = 0;
                        while (px <= tx2 && py >= ty1) {
                                if (px >= tx1 && py <= ty2) {
                                        if (maxstepheight > maxheight)
                                                maxheight = maxstepheight;
                                        ngood++;
                                        break;
                                }
                                px += vx;
                                py += vy;
                                if (py > maxstepheight)
                                        maxstepheight = py;
                                vx += vx > 0 ? -1 : (vx < 0 ? 1 : 0);
                                vy -= 1;
                        }
                }
        }

        printf("part 1: %d\n", maxheight);
        printf("part 2: %d\n", ngood);
}
