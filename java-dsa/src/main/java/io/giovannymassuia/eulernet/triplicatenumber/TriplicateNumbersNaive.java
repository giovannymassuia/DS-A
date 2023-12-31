package io.giovannymassuia.eulernet.triplicatenumber;

public class TriplicateNumbersNaive implements TriplicateNumbers {

    @Override
    public boolean isTriplicate(long number) {
        String numberString = String.valueOf(number);
//        char[] numbers = numberString.toCharArray();
//        Map<Character, Integer> counter = new HashMap<>();
//
//        for (char n : numbers) {
//            counter.merge(n, 1, Integer::sum);
//        }

        int len = numberString.length() / 3;

        for (int i = 0; i < len; i++) {
            for (int j = 0; j <= 9; j++) {
                String x = String.valueOf(j).repeat(3);
                if (numberString.contains(x)) {
                    numberString = numberString.replace(x, "");
                }

                if (numberString.isEmpty()) {
                    return true;
                }
            }
        }

        return numberString.isEmpty();
    }
}
