// https://adventofcode.com/2021/day/12

import java.util.*;

public class Day12
{
    public static void main(String[] args)
    {
        Map<String, Cave> caves = readCaves();
        Cave start = caves.get("start");
        Cave end = caves.get("end");
        System.out.printf("part 1: %d\n", findPaths(start, start, end, new HashSet<>(), false));
        System.out.printf("part 2: %d\n", findPaths(start, start, end, new HashSet<>(), true));
    }

    private static Map<String, Cave> readCaves()
    {
        Map<String, Cave> caves = new HashMap<>();
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNext()) {
            String line = scanner.next();
            if (line.isBlank())
                continue;
            String[] split = line.split("-");
            Cave a = caves.computeIfAbsent(split[0], Cave::new);
            Cave b = caves.computeIfAbsent(split[1], Cave::new);
            a.links.add(b);
            b.links.add(a);
        }
        return caves;
    }

    private static class Cave
    {
        public final String name;
        public final boolean small;
        public final List<Cave> links;

        public Cave(String name)
        {
            this.name = name;
            this.small = name.equals(name.toLowerCase());
            this.links = new ArrayList<>();
        }
    }

    private static int findPaths(Cave cave, Cave start, Cave end, Set<Cave> visited, boolean revisit)
    {
        if (cave == end)
            return 1;

        boolean second = visited.contains(cave);
        if (second && (cave == start || ! revisit))
            return 0;

        if (cave.small)
            visited.add(cave);

        int n = 0;
        for (Cave c : cave.links)
            n += findPaths(c, start, end, visited, revisit && ! second);

        if ( ! second)
            visited.remove(cave);

        return n;
    }
}
