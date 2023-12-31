package io.giovannymassuia.eulernet.triplicatenumber;

/**
 * A triplicate number is a positive integer such that, after repeatedly removing three consecutive
 * identical digits from it, all its digits can be removed.</p>
 *
 * <p>
 * For example, the integer 122555211 is a triplicate number:
 * </p>
 * <p>
 * 122[555]211 -> 1[222]11 -> [111] -> .
 * </p>
 * On the other hand, neither 663633 nor 9990 are triplicate numbers.
 */
public interface TriplicateNumbers {

    boolean isTriplicate(long number);

}
