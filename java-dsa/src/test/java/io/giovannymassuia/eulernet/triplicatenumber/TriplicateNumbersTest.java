package io.giovannymassuia.eulernet.triplicatenumber;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

class TriplicateNumbersTest {

    TriplicateNumbers triplicateNumbersNaive = new TriplicateNumbersNaive();
    TriplicateNumbers triplicateNumbersStack = new TriplicateNumbersStack();

    @ParameterizedTest
    @MethodSource("buildInputs")
    void isTriplicateNaive(long input, boolean expect) {
        assertEquals(triplicateNumbersNaive.isTriplicate(input), expect);
    }

    @ParameterizedTest
    @MethodSource("buildInputs")
    void isTriplicateStack(long input, boolean expect) {
        assertEquals(triplicateNumbersStack.isTriplicate(input), expect);
    }

    private static Stream<Arguments> buildInputs() {
        return Stream.of(
            Arguments.of(122555211L, true),
            Arguments.of(663633L, false),
            Arguments.of(9990L, false),
            Arguments.of(122555244411L, true),
            Arguments.of(122255522211L, true),
            Arguments.of(122344433211L, true),
            Arguments.of(411433314111444L, false)
        );
    }
}
