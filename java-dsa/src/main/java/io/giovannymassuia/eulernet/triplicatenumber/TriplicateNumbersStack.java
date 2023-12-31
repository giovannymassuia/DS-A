package io.giovannymassuia.eulernet.triplicatenumber;

import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

public class TriplicateNumbersStack implements TriplicateNumbers {

    @Override
    public boolean isTriplicate(long number) {
        // 122555244411L
        // [1] <-> [2] <-> [5] <-> [2] <-> [4] <-> [1]
        //  1       2       3       1       3       2
        return solutionWithTuple(number) && solutionWithMap(number);
    }

    private boolean solutionWithMap(long number) {
        Stack<Character> stack = new Stack<>();
        Map<Character, Integer> counter = new HashMap<>();

        char[] numbers = String.valueOf(number).toCharArray();

        stack.push(numbers[0]);
        counter.put(numbers[0], 1);

        for (int i = 1; i < numbers.length; i++) {
            char curr = numbers[i];

            counter.merge(curr, 1, Integer::sum);

            if (!stack.isEmpty() && stack.peek() == curr) {
                if (counter.get(curr) == 3) {
                    stack.pop();
                    counter.put(curr, 0);
                }
            } else {
                stack.push(curr);
            }
        }

        return stack.isEmpty();
    }

    private boolean solutionWithTuple(long number) {
        Stack<Tuple> stack = new Stack<>();
        char[] numbers = String.valueOf(number).toCharArray();

        for (char n : numbers) {
            if (!stack.isEmpty() && stack.peek().c == n) {
                if (++stack.peek().count == 3) {
                    stack.pop();
                }
                continue;
            }
            stack.push(Tuple.of(n));
        }

        return stack.isEmpty();
    }

    static class Tuple {

        char c;
        int count = 1;

        Tuple(char c) {
            this.c = c;
        }

        static Tuple of(char c) {
            return new Tuple(c);
        }
    }

}
