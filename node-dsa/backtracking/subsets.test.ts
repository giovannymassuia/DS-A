import {deepEqual} from "node:assert";
import test from "node:test";

/*
* Problem: Given a set of distinct integers, return all possible subsets (the power set).
* Example:
*   - Input: nums = [1,2,3]
*   - Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
* */

function subsets(nums: number[]): number[][] {
  const result: number[][] = [];

  function backtrack(start: number, path: number[]) {
    result.push([...path]); // Add the current subset to the result
    for (let i = start; i < nums.length; i++) {
      path.push(nums[i]); // Include the current element
      backtrack(i + 1, path); // Recur with the next elements
      path.pop(); // Backtrack and remove the last element
    }
  }

  backtrack(0, []); // Start backtracking from the first index
  return result;
}

test("subsets", () => {
  deepEqual(subsets([1, 2, 3]), [
    [], [1], [1, 2], [1, 2, 3], [1, 3], [2], [2, 3], [3],
  ]);
  deepEqual(subsets([0]), [[], [0]]);
  deepEqual(subsets([]), [[]]);
});
