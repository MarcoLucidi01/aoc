// https://adventofcode.com/2021/day/13

import java.util.*;
import java.util.stream.Collectors;

public class Day13
{
    private final static int X = 0;
    private final static int Y = 1;

    public static void main(String[] args)
    {
        Scanner scanner = new Scanner(System.in);
        Set<Dot> dots = readDots(scanner);
        List<Integer> folds = readFolds(scanner);

        fold(dots, folds.subList(0, 2));
        System.out.printf("part 1: %d\n", dots.size());
        fold(dots, folds.subList(2, folds.size()));
        System.out.print("part 2:\n");
        printPaper(dots);
    }

    private static Set<Dot> readDots(Scanner scanner)
    {
        Set<Dot> dots = new HashSet<>();
        while (scanner.hasNext()) {
            String line = scanner.next();
            if (line.equals("fold"))
                break;
            String[] split = line.split(",");
            dots.add(new Dot().set(Integer.parseInt(split[0]), Integer.parseInt(split[1])));
        }
        return dots;
    }

    private static List<Integer> readFolds(Scanner scanner)
    {
        List<Integer> folds = new ArrayList<>();
        while (scanner.hasNext()) {
            if ( ! scanner.next().equals("along"))
                continue;
            String inst = scanner.next();
            folds.add(inst.charAt(0) == 'x' ? X : Y);
            folds.add(Integer.parseInt(inst.substring(2)));
        }
        return folds;
    }

    private static void fold(Set<Dot> dots, List<Integer> folds)
    {
        for (int i = 0; i < folds.size(); i += 2) {
            int a = folds.get(i);
            int mid = folds.get(i + 1);
            dots
                .stream()
                .filter(d -> a == X ? d.x > mid : d.y > mid)
                .collect(Collectors.toList())
                    .stream()
                    .peek(dots::remove)
                    .map(d -> a == X ? d.set(mid * 2 - d.x, d.y) : d.set(d.x, mid * 2 - d.y))
                    .forEach(dots::add);
        }
    }

    private static void printPaper(Set<Dot> dots)
    {
        int w = maxPlusOne(dots, X);
        int h = maxPlusOne(dots, Y);
        Dot d = new Dot();
        for (int y = 0; y < h; y++) {
            for (int x = 0; x < w; x++)
                System.out.printf("%c", dots.contains(d.set(x, y)) ? '#' : '.');
            System.out.println();
        }
    }

    private static int maxPlusOne(Set<Dot> dots, int a)
    {
        return 1 + dots
            .stream()
            .map(d -> a == X ? d.x : d.y)
            .max(Comparator.comparingInt(n -> n))
            .orElse(0);
    }

    private static class Dot
    {
        public int x, y;

        public Dot set(int x, int y)
        {
            this.x = x;
            this.y = y;
            return this;
        }

        @Override
        public boolean equals(Object o)
        {
            if (this == o)
                return true;
            if (o == null || getClass() != o.getClass())
                return false;
            Dot dot = (Dot)o;
            return x == dot.x && y == dot.y;
        }

        @Override
        public int hashCode()
        {
            return Objects.hash(x, y);
        }
    }
}