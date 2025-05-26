import {deepEqual} from "node:assert";
import test from "node:test";

/*
* Problem: Given a collection of distinct integers, return all possible permutations.
* Example:
*  - Input: nums = [1,2,3]
*  - Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 */

function permute(nums: number[]): number[][] {
  const result: number[][] = [];

  function backtrack(start: number) {
    if (start === nums.length) {
      result.push([...nums]); // Add a copy of the current permutation to the result
      return;
    }
    for (let i = start; i < nums.length; i++) {
      [nums[start], nums[i]] = [nums[i], nums[start]]; // Swap the current element with the start element
      backtrack(start + 1); // Recur for the next index
      [nums[start], nums[i]] = [nums[i], nums[start]]; // Backtrack (swap back)
    }
  }

  backtrack(0); // Start backtracking from the first index
  return result;
}

test("permute", () => {
  deepEqual(permute([1, 2, 3]), [
    [1, 2, 3],
    [1, 3, 2],
    [2, 1, 3],
    [2, 3, 1],
    [3, 2, 1],
    [3, 1, 2],
  ]);
  deepEqual(permute([0]), [[0]]);
  deepEqual(permute([]), [[]]);
});
