// https://adventofcode.com/2020/day/7

import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Day7
{
    public static void main(String[] args)
    {
        RuleParser parser = new RuleParser();
        Map<String, Map<String, Integer>> rules = parser.parse(System.in);
        System.out.printf("part 1: %d\n", partOne(rules, "shiny gold"));
        System.out.printf("part 2: %d\n", partTwo(rules, "shiny gold"));
    }

    private static int partOne(Map<String, Map<String, Integer>> rules, String target)
    {
        int cnt = 0;
        for (String bagName : rules.keySet())
            if (canContain(rules, bagName, target))
                cnt++;

        return cnt;
    }

    private static boolean canContain(Map<String, Map<String, Integer>> rules, String bagName, String target)
    {
        Map<String, Integer> contain = rules.get(bagName);
        if (contain == null || contain.isEmpty())
            return false;
        if (contain.containsKey(target) && contain.get(target) > 0)
            return true;

        for (String name : contain.keySet())
            if (canContain(rules, name, target))
                return true;

        return false;
    }

    private static int partTwo(Map<String, Map<String, Integer>> rules, String target)
    {
        Map<String, Integer> contain = rules.get(target);
        if (contain == null || contain.isEmpty())
            return 0;

        int cnt = 0;
        for (String bagName : contain.keySet()) {
            int n = contain.get(bagName);
            cnt += n + n * partTwo(rules, bagName);
        }
        return cnt;
    }

    private static class RuleParser
    {
        private Scanner scanner;
        private Map<String, Map<String, Integer>> bags;

        public Map<String, Map<String, Integer>> parse(InputStream in)
        {
            scanner = new Scanner(in);
            bags = new HashMap<>();
            while (scanner.hasNext())
                rule();
            return bags;
        }

        private void accept(String want)
        {
            String got = acceptAny();
            if ( ! want.equals(got))
                throw new RuntimeException(String.format("unexpected token: want \"%s\" got \"%s\"", want, got));
        }

        private String acceptAny()
        {
            if ( ! scanner.hasNext())
                throw new RuntimeException("unexpected EOF");
            return scanner.next();
        }

        private void rule()
        {
            String bagName = bagName();
            if ( ! bags.containsKey(bagName))
                bags.put(bagName, new HashMap<>());

            accept("bags");
            accept("contain");

            if ( ! scanner.hasNextInt()) {
                accept("no");
                accept("other");
                accept("bags.");
                return;
            }

            Map<String, Integer> contain = bags.get(bagName);
            while (scanner.hasNextInt()) {
                int number = scanner.nextInt();
                String name = bagName();
                contain.put(name, number);
                acceptAny();
            }
        }

        private String bagName()
        {
            String adj = acceptAny();
            String color = acceptAny();
            return String.format("%s %s", adj, color).toLowerCase();
        }
    }
}
