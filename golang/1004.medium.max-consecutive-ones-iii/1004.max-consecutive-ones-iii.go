/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 *
 * https://leetcode.com/problems/max-consecutive-ones-iii/description/
 *
 * algorithms
 * Medium (60.90%)
 * Likes:    2185
 * Dislikes: 34
 * Total Accepted:    103.8K
 * Total Submissions: 169.8K
 * Testcase Example:  '[1,1,1,0,0,0,1,1,1,1,0]\n2'
 *
 * Given a binary array nums and an integer k, return the maximum number of
 * consecutive 1's in the array if you can flip at most k 0's.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
 * Output: 6
 * Explanation: [1,1,1,0,0,1,1,1,1,1,1]
 * Bolded numbers were flipped from 0 to 1. The longest subarray is
 * underlined.
 *
 * Example 2:
 *
 *
 * Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
 * Output: 10
 * Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
 * Bolded numbers were flipped from 0 to 1. The longest subarray is
 * underlined.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * nums[i] is either 0 or 1.
 * 0 <= k <= nums.length
 *
 *
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	/*
		用滑动窗口做。
		在窗口内，0 的个数应该<=k
		快慢指针都在开头，快指针向后走，计数遇到的0的个数
		如果 > k 了，就记录当前的长度，并移动慢指针，直到0的个数<=k
	*/
	slow, fast := 0, 0
	zeros := 0
	res := 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			zeros++
		}
		for zeros > k {
			// 右移 slow，直到 zeros == k
			// 此时 fast-slow 就是当前窗口大小
			if nums[slow] == 0 {
				zeros--
			}
			slow++
		}
		if l := fast - slow + 1; l > res {
			res = l
		}
		fast++
	}
	return res
}

// @lc code=end

