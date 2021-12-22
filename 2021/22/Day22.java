// https://adventofcode.com/2021/day/22

import java.util.*;

public class Day22
{
    public static void main(String[] args)
    {
        List<RebootStep> steps = readRebootSteps();
        System.out.printf("part 1: %d\n", partOne(steps));
        System.out.printf("part 2: %d\n", partTwo(steps));
    }

    private static List<RebootStep> readRebootSteps()
    {
        List<RebootStep> steps = new ArrayList<>();
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNext()) {
            boolean on = scanner.next().equals("on");
            String[] range = scanner.next().split(",");
            String[] x = range[0].replace("x=", "").split("\\.\\.");
            String[] y = range[1].replace("y=", "").split("\\.\\.");
            String[] z = range[2].replace("z=", "").split("\\.\\.");
            steps.add(new RebootStep(on, new Cuboid(
                Integer.parseInt(x[0]), Integer.parseInt(x[1]),
                Integer.parseInt(y[0]), Integer.parseInt(y[1]),
                Integer.parseInt(z[0]), Integer.parseInt(z[1])
            )));
        }
        return steps;
    }

    private static class RebootStep
    {
        public final boolean on;
        public final Cuboid cuboid;

        public RebootStep(boolean on, Cuboid cuboid)
        {
            this.on = on;
            this.cuboid = cuboid;
        }
    }

    private static class Cuboid
    {
        public final int xfrom, xto;
        public final int yfrom, yto;
        public final int zfrom, zto;

        public Cuboid(int xfrom, int xto, int yfrom, int yto, int zfrom, int zto)
        {
            this.xfrom = xfrom;
            this.xto = xto;
            this.yfrom = yfrom;
            this.yto = yto;
            this.zfrom = zfrom;
            this.zto = zto;
        }

        public long volume()
        {
            return (long) (xto - xfrom + 1) * (yto - yfrom + 1) * (zto - zfrom + 1);
        }
    }

    private static int partOne(List<RebootStep> steps)
    {
        // init procedure only uses cubes that have x,y,z positions of at least -50 and at most 50
        final int N = 51;

        // this obviously won't work for part two
        boolean[][][] states = new boolean[N*2][N*2][N*2];
        for (RebootStep step : steps)
            for (int x = Math.max(-N, step.cuboid.xfrom); x <= Math.min(step.cuboid.xto, N); x++)
                for (int y = Math.max(-N, step.cuboid.yfrom); y <= Math.min(step.cuboid.yto, N); y++)
                    for (int z = Math.max(-N, step.cuboid.zfrom); z <= Math.min(step.cuboid.zto, N); z++)
                        states[x+N][y+N][z+N] = step.on;

        int cnt = 0;
        for (boolean[][] a : states)
            for (boolean[] b : a)
                for (boolean on : b)
                    if (on)
                        cnt++;

        return cnt;
    }

    private static long partTwo(List<RebootStep> steps)
    {
        // TODO
        return 0;
    }
}
