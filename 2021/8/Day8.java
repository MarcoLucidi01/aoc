// https://adventofcode.com/2021/day/8

import java.io.InputStream;
import java.util.*;

public class Day8
{
    public static void main(String[] args)
    {
        List<Entry> entries = readEntries(System.in);
        System.out.printf("part 1: %d\n", partOne(entries));
        System.out.printf("part 2: %d\n", partTwo(entries));
    }

    private static List<Entry> readEntries(InputStream in)
    {
        List<Entry> entries = new ArrayList<>();
        Scanner scanner = new Scanner(in);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            if (line.isBlank())
                continue;
            String[] lineSplit = line.split("\\s+\\|\\s+");
            String[] patterns = lineSplit[0].split("\\s+");
            String[] outputs = lineSplit[1].split("\\s+");
            entries.add(new Entry(patterns, outputs));
        }
        return entries;
    }

    private static class Entry
    {
        public final String[] patterns;
        public final String[] outputs;

        public Entry(String[] patterns, String[] outputs)
        {
            this.patterns = patterns;
            this.outputs = outputs;
        }
    }

    private static long partOne(List<Entry> entries)
    {
        Set<Integer> uniqueLengths = Set.of(2, 3, 4, 7);
        return entries
            .stream()
            .flatMap(entry -> Arrays.stream(entry.outputs))
            .filter(output -> uniqueLengths.contains(output.length()))
            .count();
    }

    private static int partTwo(List<Entry> entries)
    {
        int sum = 0;
        var decoder = new SevenSegmentsDecoder();
        for (var entry : entries) {
            decoder.decode(entry.patterns);
            int n = 0;
            for (var output : entry.outputs)
                n = n * 10 + decoder.getDigit(output);
            sum += n;
        }
        return sum;
    }

    private static class SevenSegmentsDecoder
    {
        private Map<String, Integer> digitsTable;   // pattern -> digit
        private Map<Integer, String> patternsTable; // digit -> pattern

        public void decode(String[] patterns)
        {
            clear();
            Arrays.sort(patterns, Comparator.comparingInt(String::length));
            for (String pattern : patterns) {
                switch (pattern.length()) {
                case 2:
                    put(pattern, 1);
                    break;
                case 3:
                    put(pattern, 7);
                    break;
                case 4:
                    put(pattern, 4);
                    break;
                case 5:
                    if (countOverlappingSegments(pattern, patternsTable.get(1)) == 2)
                        put(pattern, 3);
                    else if (countOverlappingSegments(pattern, patternsTable.get(4)) == 3)
                        put(pattern, 5);
                    else
                        put(pattern, 2);
                    break;
                case 6:
                    if (countOverlappingSegments(pattern, patternsTable.get(4)) == 4)
                        put(pattern, 9);
                    else if (countOverlappingSegments(pattern, patternsTable.get(7)) == 3)
                        put(pattern, 0);
                    else
                        put(pattern, 6);
                    break;
                case 7:
                    put(pattern, 8);
                    break;
                }
            }
        }

        private void clear()
        {
            digitsTable = new HashMap<>();
            patternsTable = new HashMap<>();
        }

        private void put(String pattern, int digit)
        {
            pattern = normalize(pattern);
            digitsTable.put(pattern, digit);
            patternsTable.put(digit, pattern);
        }

        public int getDigit(String pattern)
        {
            return digitsTable.get(normalize(pattern));
        }

        private static String normalize(String pattern)
        {
            char[] chars = pattern.toCharArray();
            Arrays.sort(chars);
            return new String(chars);
        }

        private static long countOverlappingSegments(String pattern1, String pattern2)
        {
            return pattern1
                .chars()
                .filter(c -> pattern2.indexOf(c) >= 0)
                .count();
        }
    }
}
