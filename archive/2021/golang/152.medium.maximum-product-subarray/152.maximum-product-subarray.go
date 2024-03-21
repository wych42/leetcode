/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (32.73%)
 * Likes:    6560
 * Dislikes: 213
 * Total Accepted:    465.1K
 * Total Submissions: 1.4M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find a contiguous non-empty subarray within the
 * array that has the largest product, and return the product.
 *
 * It is guaranteed that the answer will fit in a 32-bit integer.
 *
 * A subarray is a contiguous subsequence of the array.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 因为乘积可能出现负数，负数在后面有可能变成正的，也可能不会, 所以要记录下来
	// 在走到 i 的时候:
	// 正数有三个选项: i 自己；上一个正数*i，上一个负数*i 里最大的那个
	// 负数有三个选项: i 自己；上一个正数*i，上一个负数*i 里最小的那个
	dpmax := make([]int, len(nums))
	dpmin := make([]int, len(nums))
	dpmax[0] = nums[0]
	dpmin[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		// nums[i]
		// nums[i]*dpmax[i-1]
		// nums[i]*dpmin[i-1]
		dpmax[i] = max(n, max(dpmax[i-1]*n, dpmin[i-1]*n))
		dpmin[i] = min(n, min(dpmin[i-1]*n, dpmax[i-1]*n))
	}
	//	fmt.Println("dpmax", dpmax)
	//	fmt.Println("dpmin", dpmin)
	ans := 0
	for _, n := range dpmax {
		if n > ans {
			ans = n
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

