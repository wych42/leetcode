package main

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.cn/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (75.28%)
 * Likes:    1750
 * Dislikes: 0
 * Total Accepted:    403.1K
 * Total Submissions: 535.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)

 * brute force:  O(n^2), loop with 2-depth
 * O(n) with division: loop 1: multiple all, loop 2: divide each
 * O(n) without division:
   for index i, result is computed by numbers before i, and numbers after i
   use a array to store incremental product result.
   [1,2,3,4]
   for 1, product by Nan and 2,3,4
   for 2, product by 1 and 3,4
   for 3, product by 1,2 and 4
   for 4, product by 1,2,3 and Nan
 * O(n) in time and O(1) in space: product prefix and store in result, then product suffix

*/

// @lc code=start

// O1 space
func productExceptSelf(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	// left to right
	// input 1,2,3,4
	// result 1, 0, 0,0
	//       1, 1*1, 0, 0
	//       1, 1, 2, 0
	//       1, 1, 2, 6
	result[0] = 1
	for i := 1; i < size; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	// right to left
	suffix := 1 // track suffix icremental product
	for i := size - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}

func productExceptSelf_OnSpace(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	prefix := make([]int, size)
	prefix[0] = 1
	suffix := make([]int, size)
	suffix[size-1] = 1

	for i := 1; i < size; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := size - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i := 0; i < size; i++ {
		result[i] = prefix[i] * suffix[i]
	}
	return result
}

// @lc code=end
