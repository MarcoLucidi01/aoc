// https://adventofcode.com/2021/day/18

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Day18
{
    public static void main(String[] args)
    {
        List<String> numbers = readNumbers();
        System.out.printf("part 1: %d\n", partOne(numbers));
        System.out.printf("part 2: %d\n", partTwo(numbers));
    }

    private static List<String> readNumbers()
    {
        List<String> numbers = new ArrayList<>();
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String n = scanner.nextLine();
            if ( ! n.isBlank())
                numbers.add(n);
        }
        return numbers;
    }

    private static int partOne(List<String> numbers)
    {
        if (numbers.isEmpty())
            return 0;

        SnailNumber sum = SnailNumber.parse(numbers.get(0));
        for (int i = 1; i < numbers.size(); i++)
            sum = sum.add(SnailNumber.parse(numbers.get(i)));

        return sum.magnitude();
    }

    private static int partTwo(List<String> numbers)
    {
        int maxMagnitude = 0;
        for (String a : numbers) {
            for (String b : numbers) {
                if (a.equals(b))
                    continue;
                // SnailNumber(s) are not immutable, that's why I have to parse() every time.
                int m = SnailNumber.parse(a).add(SnailNumber.parse(b)).magnitude();
                if (m > maxMagnitude)
                    maxMagnitude = m;
            }
        }
        return maxMagnitude;
    }

    private static class SnailNumber
    {
        public SnailNumber parent;
        public SnailNumber left;
        public SnailNumber right;
        public Integer value;

        private SnailNumber(SnailNumber parent, SnailNumber left, SnailNumber right, Integer value)
        {
            this.parent = parent;
            this.left = left;
            this.right = right;
            this.value = value;
        }

        public static SnailNumber parse(String s)
        {
            return parse(new Scanner(s).useDelimiter("")); // scan char by char.
        }

        private static SnailNumber parse(Scanner scanner)
        {
            SnailNumber sn = new SnailNumber(null, null, null, null);

            String c = scanner.next();
            if (c.equals("[")) {
                sn.left = parse(scanner);
                sn.left.parent = sn;
                c = scanner.next();
                if ( ! c.equals(","))
                    throw new RuntimeException(String.format("want ',' but got '%s'", c));

                sn.right = parse(scanner);
                sn.right.parent = sn;
                c = scanner.next();
                if ( ! c.equals("]"))
                    throw new RuntimeException(String.format("want ']' but got '%s'", c));

                return sn;
            }

            sn.value = Integer.valueOf(c);
            return sn;
        }

        public SnailNumber add(SnailNumber sn)
        {
            parent = new SnailNumber(null, this, sn, null);
            sn.parent = parent;
            parent.reduce();
            return parent;
        }

        public void reduce()
        {
            while (explode(0) || split())
                ;
        }

        private boolean explode(int level)
        {
            if (isRegular())
                return false;

            if (level < 4)
                return left.explode(level + 1) || right.explode(level + 1);

            if ( ! left.isRegular() || ! right.isRegular())
                return false;

            addLeft();
            addRight();
            left = null;
            right = null;
            value = 0;
            return true;
        }

        private void addLeft()
        {
            SnailNumber par = parent;
            SnailNumber cur = this;
            while (par != null && cur == par.left) { // go up
                par = par.parent;
                cur = cur.parent;
            }
            if (par == null)
                return;

            cur = par.left;
            while ( ! cur.isRegular()) // go down
                cur = cur.right;

            cur.value += left.value;
        }

        private void addRight()
        {
            SnailNumber par = parent;
            SnailNumber cur = this;
            while (par != null && cur == par.right) { // go up
                par = par.parent;
                cur = cur.parent;
            }
            if (par == null)
                return;

            cur = par.right;
            while ( ! cur.isRegular()) // go down
                cur = cur.left;

            cur.value += right.value;
        }

        private boolean split()
        {
            if ( ! isRegular())
                return left.split() || right.split();

            if (value < 10)
                return false;

            left = new SnailNumber(this, null, null, Math.floorDiv(value, 2));
            right = new SnailNumber(this, null, null, Math.floorDiv(value + 1, 2));
            value = null;
            return true;
        }

        public boolean isRegular()
        {
            return value != null && left == null && right == null;
        }

        public int magnitude()
        {
            if (isRegular())
                return value;

            return 3 * left.magnitude() + 2 * right.magnitude();
        }

        @Override
        public String toString()
        {
            if (isRegular())
                return value.toString();

            return String.format("[%s,%s]", left.toString(), right.toString());
        }
    }
}
