// https://adventofcode.com/2021/day/7

import java.io.InputStream;
import java.util.*;
import java.util.function.Function;
import java.util.stream.Collectors;

public class Day7
{
    public static void main(String[] args)
    {
        List<Integer> pos = readPositions(System.in);
        System.out.printf("part 1: %d\n", minimumAlignFuel(pos, dist -> dist));
        System.out.printf("part 2: %d\n", minimumAlignFuel(pos, dist -> (dist * (dist + 1)) / 2));
    }

    public static List<Integer> readPositions(InputStream in)
    {
        return Arrays
            .stream(new Scanner(in)
                .nextLine()
                .split(","))
            .map(String::trim)
            .map(Integer::valueOf)
            .collect(Collectors.toList());
    }

    public static int minimumAlignFuel(List<Integer> positions, Function<Integer, Integer> fuelCalculator)
    {
        int maxPos = Collections.max(positions);
        int minPos = Collections.min(positions);
        int minFuel = Integer.MAX_VALUE;
        for (int end = minPos; end <= maxPos; end++) {
            int fuel = 0;
            for (int start : positions)
                fuel += fuelCalculator.apply(Math.abs(end - start));
            if (fuel < minFuel)
                minFuel = fuel;
        }
        return minFuel;
    }
}
