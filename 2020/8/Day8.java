// https://adventofcode.com/2020/day/8

import java.io.InputStream;
import java.util.*;

public class Day8
{
    public static void main(String[] args)
    {
        Instruction[] program = readProgram(System.in);
        System.out.printf("part 1: %d\n", partOne(program));
        System.out.printf("part 2: %d\n", partTwo(program));
    }

    private enum Operation
    {
        ACC,
        JMP,
        NOP,
    }

    private static class Instruction
    {
        public Operation op;
        public int arg;

        public Instruction(Operation op, int arg)
        {
            this.op = op;
            this.arg = arg;
        }
    }

    private static Instruction[] readProgram(InputStream in)
    {
        List<Instruction> program = new ArrayList<>();
        Scanner scanner = new Scanner(in);
        while (scanner.hasNext()) {
            Operation op = Operation.valueOf(scanner.next().toUpperCase());
            int arg = 0;
            if (scanner.hasNextInt())
                arg = scanner.nextInt();
            program.add(new Instruction(op, arg));
        }
        return program.toArray(new Instruction[0]);
    }

    private static int partOne(Instruction[] program)
    {
        VM vm = new VM();
        vm.runTillEnd(program);
        return vm.accumulator;
    }

    private static int partTwo(Instruction[] program)
    {
        VM vm = new VM();
        for (Instruction inst : program) {
            switch (inst.op) {
            case JMP:
                inst.op = Operation.NOP;
                if (vm.runTillEnd(program))
                    return vm.accumulator;
                inst.op = Operation.JMP;
                break;
            case NOP:
                inst.op = Operation.JMP;
                if (vm.runTillEnd(program))
                    return vm.accumulator;
                inst.op = Operation.NOP;
                break;
            }
        }
        return vm.accumulator;
    }

    private static class VM
    {
        public int accumulator;

        public boolean runTillEnd(Instruction[] program)
        {
            accumulator = 0;
            Set<Instruction> executed = new HashSet<>();
            for (int pc = 0; pc < program.length; ) {
                Instruction inst = program[pc];
                if (executed.contains(inst))
                    return false;
                switch (inst.op) {
                case ACC:
                    accumulator += inst.arg;
                    pc++;
                    break;
                case JMP:
                    pc += inst.arg;
                    break;
                case NOP:
                    pc++;
                    break;
                }
                executed.add(inst);
            }
            return true;
        }
    }
}
